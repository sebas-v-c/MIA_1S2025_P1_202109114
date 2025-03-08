import './App.css'
import Topbar from "./components/Topbar.tsx";
import Console from "./components/Console.tsx";
import Output from "./components/Output.tsx";
import {useState} from "react";
import axios, {AxiosResponse} from "axios";

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
    line: number,
    column: number,
    msg: string
}

export type RunError = {
    line: number,
    column: number,
    command: Command,
    msg: string
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

    const handleRun = async() => {
        setSynErrors([])
        setRunErrors([])
        setCommandLogs(["Running..."]);

        axios.post<{code: string}, AxiosResponse<ApiResponse>>("http://localhost:8080/parse", {"code": code}).then((response) =>{
            setSynErrors(response.data.synErrors);
            setRunErrors(response.data.runErrors);
            setCommandLogs(response.data.commandLogs);
            console.log(response);
        }).catch(error => {
            console.log(error)
        })
    }


    return (
        <div className="h-screen flex flex-col bg-[#1e1e2e] text-white ">
            <Topbar onFileUpload={setCode} onRun={handleRun}/>
            <Console code={code} onCodeChange={setCode}/>
            <Output synErrors={synErrors} runErrors={runErrors} commandLogs={commandLogs}/>
        </div>
    )
}

export default App
