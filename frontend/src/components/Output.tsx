import {RunError, SynError} from "../App.tsx";

interface OutputProps {
    synErrors: Array<SynError>,
    runErrors: Array<RunError>,
    commandLogs: Array<string>
}

function Output({synErrors, runErrors, commandLogs}: OutputProps) {
    return (
        <div className="w-full h-[100%] bg-[#282a36] text-[#f8f8f2] p-4 overflow-auto">
            <pre className="whitespace-pre-wrap">
                {/* Display syntax errors in yellow */}
                { synErrors.length > 0 && (
                    <div className="text-yellow-400">
                        <strong>Syntax Errors:</strong>
                        {
                            synErrors.map((err, index) => (
                                <div key={index} className="mb-5">
                                    [Line {err.Line}, Col {err.Column}]: {err.Msg}
                                </div>
                            ))
                        }
                    </div>
                )}
                {/* Display runtime errors in red */}
                { runErrors.length > 0 && (
                    <div className="text-red-400">
                        <strong>Runtime Errors:</strong>
                        {runErrors.map((err, index)=>(
                            <div key={index} className="mb-5">
                                [Line {err.Line}, Col {err.Column}, Command {err.Command}]: {err.Msg}
                            </div>
                        ))}
                    </div>
                )}
                {/* Display command logs in green */}
                { commandLogs.length > 0 && (
                    <div className="text-green-400">
                        <strong>Execution Logs:</strong>
                        {commandLogs.map((log, index)=>(
                            <div key={index} className="mb-5">
                                $ {log}
                            </div>
                        ))}
                    </div>
                )}
            </pre>
        </div>
    );

}

export default Output