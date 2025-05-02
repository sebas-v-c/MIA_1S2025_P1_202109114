import './App.css'
import Topbar from "./components/Topbar.tsx";
import Console from "./components/Console.tsx";
import Output from "./components/Output.tsx";
import {useState} from "react";
import axios, {AxiosResponse} from "axios";
import toast, {Toaster} from "react-hot-toast";
import * as React from "react";
import Explorer from "./components/explorer/Explorer.tsx";

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

function App() {
    const [code, setCode] = useState<string>("# Write some code here...")
    const [synErrors, setSynErrors] = useState<Array<SynError>>([])
    const [runErrors, setRunErrors] = useState<Array<RunError>>([])
    const [commandLogs, setCommandLogs] = useState<Array<string>>([])
    const [consoleMode, setConsoleMode] = React.useState(false);

    const handleRun = async() => {
        setSynErrors([])
        setRunErrors([])
        setCommandLogs(["Running..."]);

        axios.post<{code: string}, AxiosResponse<ApiResponse>>("http://localhost:8080/parse", {"code": code? code : " "}).then((response) =>{
            console.log(response.data.runErrors)
            setSynErrors(response.data.synErrors? response.data.synErrors : []);
            setRunErrors(response.data.runErrors? response.data.runErrors : []);
            setCommandLogs(response.data.commandLogs? response.data.commandLogs : []);
            console.log(response);
        }).catch(error => {
            toast.error("Failed to connect to the API, Check internet or server")
            console.log(error)
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


    return (
        <div className="h-screen flex flex-col bg-[#1e1e2e] text-white ">
            <Toaster position="bottom-right" reverseOrder={false} />
            <Topbar onFileUpload={setCode} onRun={handleRun} onClear={handleClear} onExplorer={handleExplore} barMode={consoleMode}/>
            {consoleMode? (
                <>
                    <Console code={code} onCodeChange={setCode}/>
                    <Output synErrors={synErrors} runErrors={runErrors} commandLogs={commandLogs}/>
                </>
            ) : (
                <>
                    <Explorer/>
                </>
            )}
        </div>
    )
}

export default App
