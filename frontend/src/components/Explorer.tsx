// Explorer.tsx
import {useEffect, useState} from "react";
import axios from "axios";
import FileExplorer, {File, FileType} from "./explorer/FileExplorer";
import {LoggedUser, MountedPartition, User} from "../App.tsx";
import LogIn from "./explorer/LogIn.tsx";
import toast from "react-hot-toast";

// Types matching backend DiscInfo & PartitionInfo JSON
export type PartitionInfo = {
    status: string;
    type: string;
    fit: string;
    start: number;
    size: number;
    name: string;
    correlative: number;
    id: string;
    partType: string
};

export type DiscInfo = {
    name: string;            // Ensure backend includes `name` field for each disc
    mbrSize: number;
    creationDate: string;
    signature: number;
    fit: string;
    partitions: PartitionInfo[];
};

type ApiResponse = {
    type: "list_discs";
    contents: DiscInfo[];
};

type FileRes = {
    name: string,
    type: string,
    size: number,
    atime: string,
    mtime: string,
    ctime: string,
    owner: string,
    group: string,
    permissions: string,
    content?: string,
}

type ApiFileResponse = {
    error?: string;
    files?: FileRes[];
}

interface ExplorerProps {
    onLogin: (user: User, parition: MountedPartition) => void,
    currentUser?: LoggedUser | null
}

export default function Explorer({onLogin, currentUser}: ExplorerProps) {
    const [files, setFiles] = useState<File[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);
    const [currentPath, setCurrentPath] = useState<string>("%");
    const [showLogIn, setShowLogIn] = useState<boolean>(false);
    const [defaultIdPartition, setDefaultIdPartition] = useState<string>("");

    const fetchDiscs = async (path: string) => {
        setLoading(true);
        setError(null);
        try {
            const res = await axios.post<ApiResponse>(
                "http://ec2-18-218-94-128.us-east-2.compute.amazonaws.com:80/discs",
                {path}
            );

            // Map DiscInfo[] → File[] for FileExplorer
            const discFiles: File[] = res.data.contents.map(d => ({
                name: d.name,
                size: d.mbrSize,
                mbrSize: d.mbrSize,
                type: "Disc" as FileType,
                createdTime: d.creationDate,
                modifiedTime: d.creationDate,
                signature: d.signature,
                fit: d.fit,
                content: undefined,
                // Embed partitions for drill-down
                partitions: d.partitions.map(p => ({
                    name: p.name,
                    size: p.size,
                    type: "Partition" as FileType,
                    createdTime: d.creationDate,
                    modifiedTime: d.creationDate,
                    content: undefined,
                    groupId: 0,
                    userId: 0,
                    status: p.status,
                    fit: p.fit,
                    start: p.start,
                    correlative: p.correlative,
                    id: p.id,
                    partType: p.type,
                })),
            }));
            console.log(discFiles)

            setFiles(discFiles);
        } catch (err: unknown) {
            setError(
                err instanceof Error ? err.message : "Failed to fetch discs"
            );
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        fetchDiscs("%");
    }, []);

    if (loading) return <div className="text-white">Loading...</div>;
    if (error) return <div className="text-red-500">Error: {error}</div>;

    const handleFileDoubleClick = (disc: File) => {
        if (disc.type === "Disc") {
            const partitionFiles: File[] = disc.partitions
                ? disc.partitions.map(partition => ({
                    name: partition.name,
                    size: partition.size,
                    type: partition.type,
                    createdTime: partition.createdTime,
                    modifiedTime: partition.modifiedTime,
                    content: undefined,
                    status: partition.status,
                    fit: partition.fit,
                    start: partition.start,
                    correlative: partition.correlative,
                    id: partition.id,
                    partType: partition.partType,
                }))
                : [];
            setFiles(partitionFiles);
            setCurrentPath(currentPath + disc.name)
            return
        }

        if (disc.type === "Partition") {
          if (!currentUser){
            /* MAKE THE LOG IN COMPONENT VISIBLE IN WHOLE SCREEN*/
            setShowLogIn(true);
            setDefaultIdPartition(disc.id || "")
          }
          if(currentUser){

              //@ts-expect-error some bad parameters here
              if (disc.id !== String.fromCharCode(...currentUser.MountedPartition.Id)){
                  toast.error("You can't log in into this partition because you are already logged in into another partition. Please log out from the current partition first.")
                  return
              }
              setCurrentPath(currentPath + "%" + disc.name + "%" )
              requestExplore("/")
          }
        }
    }

    const hanldeFolderDoubleClick = (folder: File) => {
        if (currentUser){
            const paths = currentPath.split("%")
            const filePath = paths[paths.length - 1] + "/" +folder.name
            filePath.slice(1)

            console.log(filePath)
            requestExplore(filePath)
        }
    }


    const navigateUp = () => {
        // 1) turn "%Disc1%Partition1%/home/zibas/users" into ["Disc1","Partition1","/home/zibas/users"]
        const segments = currentPath
            .split("%")
            .filter(part => part !== "");

        // 2) if nothing left → already at root
        if (segments.length === 0) {
            toast.custom(t => (
                <div className={
                    `bg-blue-500 text-white px-4 py-2 rounded shadow-lg ` +
                    (t.visible ? 'animate-enter' : 'animate-leave')
                }>
                    Already at the root directory.
                </div>
            ));
            return;
        }

        // 3) look at the last segment
        const last = segments[segments.length - 1];

        if (last.startsWith("/")) {
            // ── a filesystem path inside a partition: "/home/zibas/users"
            const dirs = last
                .split("/")         // ["","home","zibas","users"]
                .filter(Boolean);   // ["home","zibas","users"]

            dirs.pop();           // remove "users"

            if (dirs.length > 0) {
                // still in a sub-directory, e.g. "/home/zibas"
                segments[segments.length - 1] = "/" + dirs.join("/");
            } else {
                // popped back to "/" → drop that entire segment (back to partition root)
                segments.pop();
            }
        } else {
            // ── disc or partition level, just drop it
            segments.pop();
        }

        // 4) rebuild path; "%" means top‐level of all discs
        const newPath = segments.length > 0
            ? "%" + segments.join("%")
            : "%";

        setCurrentPath(newPath);

        // 5) trigger the right fetch
        if (newPath === "%") {
            fetchDiscs("%");
        } else {
            // last real segment = either "Partition1" or "/home/zibas"
            const parts = newPath.split("%").filter(Boolean);
            const toExplore = parts[parts.length - 1];
            requestExplore(toExplore);
        }
    };



    const handleLogIn = (user: User, partition: MountedPartition) => {
        onLogin(user, partition);
        setShowLogIn(false)
    }

const requestExplore = async (path: string) => {
    try {
        // Call the /explore endpoint with the proper body
        const response = await axios.post<ApiFileResponse>(
            "http://ec2-18-218-94-128.us-east-2.compute.amazonaws.com:80/explore",
            { "path": path }
        );

        // Process the response
        const { error, files } = response.data;

        if (error) {
            toast.error(`Error: ${error}`);
            return;
        }

        const filesToShow: File[] = files ? files.map(file => ({
            name: file.name,
            size: file.size,
            type: file.type === "0" ? "Directory" : "File" as FileType,
            createdTime: file.ctime,
            modifiedTime: file.mtime,
            accessedTime: file.atime,
            content: file.content,
            owner: file.owner,
            group: file.group,
            permissions: file.permissions,
        })) : [];

        setFiles(filesToShow);
        setCurrentPath(path);
        toast.success("Files loaded successfully!");
    } catch (err) {
        console.error(err);
        toast.error("Failed to load the directory. Please try again.");
    }
};


    return (
        <div className="h-full flex flex-col">
            {/* Path bar header */}
            <div className="bg-gray-200 text-gray-800 py-2 px-4 border-b border-gray-300">
                <span className="font-mono text-sm">{currentPath.replace(/%/g, "/") || "/"}</span>
            </div>

            {/* File explorer */}
            <div className="flex-grow">
                <FileExplorer
                    files={files}
                    onFileDoubleClick={handleFileDoubleClick}
                    onNavigateUp={navigateUp}
                    onFolderDoubleClick={hanldeFolderDoubleClick}
                />
            </div>

            {/* Log in component */}
            {showLogIn && (
                <LogIn
                    defaultIdPartition={defaultIdPartition}
                    onLogin={handleLogIn}
                    onClose={() => setShowLogIn(false)}
                />
            )}
        </div>
    );
}