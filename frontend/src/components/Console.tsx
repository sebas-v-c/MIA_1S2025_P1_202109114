import {Editor} from "@monaco-editor/react";

interface ConsoleProps {
    code: string,
    onCodeChange: (value: (((prevState: string) => string) | string)) => void
}

function Console({code, onCodeChange}: ConsoleProps) {

    return (
        <div className="w-full h-[50vh] bg-[#1e1e2e]">
            <Editor
                height="100%"
                defaultLanguage="shell"
                theme="vs-dark"
                value={code}
                onChange={(value) => onCodeChange(value ? value : "")}
                options={{
                    fontSize: 14,
                    minimap: {enabled: false},
                }}
            />
        </div>
    )
}

export default Console