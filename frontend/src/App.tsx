import './App.css'
import Topbar from "./components/Topbar.tsx";
import Console from "./components/Console.tsx";
import Output from "./components/Output.tsx";
import {useState} from "react";
import axios, {AxiosResponse} from "axios";
import toast, {Toaster} from "react-hot-toast";
import * as React from "react";
import Explorer, {PartitionInfo} from "./components/Explorer.tsx";

export enum Command {
    MKDISK,
    RMDISK,
    FDISK,
    MOUNT,
    MOUNTED,
    MKFS,
    CAT,
    LOGIN,
    LOGOUT,
    MKGROUP,
    RMGROUP,
    MKUSR,
    RMUSR,
    CHRGP,
    MKFILE,
    MKDIR,
    REP,
}

export type SynError = {
    Line: number,
    Column: number,
    Msg: string
}

export type RunError = {
    Line: number,
    Column: number,
    Command: Command,
    Msg: string
}

export type ApiResponse = {
    synErrors: Array<SynError>,
    runErrors: Array<RunError>,
    commandLogs: Array<string>
}

export type User = {
    Id: number[],
    Group: string,
    Name: string,
    Password: string
}

export type MountedPartition = PartitionInfo &{
    DiscSignature: number,
    DiscTag: string,
    Path: string
}

export type LoggedUser = {
    User: User,
    MountedPartition: MountedPartition
}

function App() {
    const [code, setCode] = useState<string>("# Write some code here...")
    const [synErrors, setSynErrors] = useState<Array<SynError>>([])
    const [runErrors, setRunErrors] = useState<Array<RunError>>([])
    const [commandLogs, setCommandLogs] = useState<Array<string>>([])
    const [consoleMode, setConsoleMode] = React.useState(false);
    const [loggedUser, setLoggedUser] = useState<LoggedUser | null>(null);

    const handleRun = async() => {
        setSynErrors([])
        setRunErrors([])
        setCommandLogs(["Running..."]);

        axios.post<{code: string}, AxiosResponse<ApiResponse>>("http://localhost:8080/parse", {"code": code? code : " "}).then((response) =>{
            setSynErrors(response.data.synErrors? response.data.synErrors : []);
            setRunErrors(response.data.runErrors? response.data.runErrors : []);
            setCommandLogs(response.data.commandLogs? response.data.commandLogs : []);
        }).catch(error => {
            console.error("API request failed:", error);
            toast.error("Failed to connect to the API, Check internet or server");
        })
    }

    const handleClear = () => {
        setSynErrors([])
        setRunErrors([])
        setCommandLogs([])
        setCode("")
    }

    const handleExplore = () => {
        setConsoleMode(!consoleMode);
    }

    const handleLogin = (user: User, partition: MountedPartition) => {
        setLoggedUser({User: user, MountedPartition: partition});
    }

    const handleLogout = async (e:React.FormEvent) => {
        e.preventDefault();
        const parseResponse = await axios.post("http://localhost:8080/parse", { code: "\nmounted\nlogout\nmounted\n" }, {
        });

        const { synErrors, runErrors } = parseResponse.data;

        // Handle syntax or runtime errors from the parse step
        if (synErrors || runErrors.length > 0) {
            if (synErrors) {
                //@ts-expect-error any implementation
                toast.error("Syntax Errors:\n" + synErrors.map((e) => e.Msg).join("\n"));
            }
            if (runErrors.length > 0) {
                //@ts-expect-error any implementation
                toast.error("Runtime Errors:\n" + runErrors.map((e) => e.Msg).join("\n"));
            }
            console.error("Parse Errors:", {
                synErrors,
                runErrors,
            });
            return;
        }

        setLoggedUser(null);
    }

    return (
        <div className="h-screen flex flex-col bg-[#1e1e2e] text-white ">
            <Toaster position="bottom-right" reverseOrder={false} />
            <Topbar onFileUpload={setCode} onRun={handleRun} onClear={handleClear} onExplorer={handleExplore} barMode={consoleMode} loggedUser={loggedUser} onLogOut={handleLogout}/>
            {consoleMode? (
                <>
                    <Console code={code} onCodeChange={setCode}/>
                    <Output synErrors={synErrors} runErrors={runErrors} commandLogs={commandLogs}/>
                </>
            ) : (
                <>
                    <Explorer onLogin={handleLogin} currentUser={loggedUser}/>
                </>
            )}
        </div>
    )
}

export default App
