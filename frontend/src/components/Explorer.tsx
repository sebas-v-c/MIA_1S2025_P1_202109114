// Explorer.tsx
import {useEffect, useState} from "react";
import axios from "axios";
import FileExplorer, {File, FileType} from "./explorer/FileExplorer";
import {LoggedUser, MountedPartition, User} from "../App.tsx";
import LogIn from "./explorer/LogIn.tsx";

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
                "http://localhost:8080/discs",
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
              return;
          }
          if(currentUser){
              setCurrentPath(currentPath + "%" + disc.name + "/")

          }


        }
    }

    const navigateUp = () => {
        let path = currentPath.split("%")
        let newPath = ""
        if (path.length > 0) {
            path = path.slice(0, path.length - 1)

            newPath = path.join("%")
            newPath = "%" + newPath
        }

        if (newPath === "%") {
            fetchDiscs("%")
        }
    }

    const handleLogIn = (user: User, partition: MountedPartition) => {
        onLogin(user, partition);
        setShowLogIn(false)
    }


    return (
        <div className="h-full">
            <FileExplorer
                files={files}
                onFileDoubleClick={handleFileDoubleClick}
                onNavigateUp={navigateUp}
            />
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
