import * as React from "react";

interface TopbarProps {
    onFileUpload: (value: (((prevState: string) => string) | string)) => void,
    onRun: () => Promise<void>,
    onClear: () => void
}

function Topbar({onFileUpload, onRun, onClear}: TopbarProps) {

    function handleFileUpload(event: React.ChangeEvent<HTMLInputElement>) {
        const file = event.target.files?.[0];
        if (file) {
            const reader = new FileReader();

            reader.onload = (e) => {
                const text = e.target?.result as string;
                onFileUpload(text)
            };
            reader.readAsText(file);
        }

    }

    return (
        <nav className="bg-[#282a36] text-white p-4 shadow-lg flex justify-between items-center">
            <h1 className="text-xl font-bold">EXT2 sim</h1>
            <div className="space-x-4 flex items-center">
                <button
                    onClick={onRun}
                    className="bg-[#58f88e] hover:bg-[#ff79c6] px-4 py-2 rounded-lg cursor-pointer transition">
                    Run
                </button>
                <input
                    type="file"
                    accept=".smia"
                    className="hidden"
                    id="fileUpload"
                    onChange={handleFileUpload}
                />
                <label
                    htmlFor="fileUpload"
                    className="bg-[#ffb86c] hover:bg-[#ff79c6] px-4 py-2 rounded-lg cursor-pointer transition"
                >
                    Upload File
                </label>
                <button
                    className="bg-[#ff5555] hover:bg-[#ff79c6] px-4 py-2 rounded-lg cursor-pointer transition"
                    onClick={onClear}
                >
                    Clear
                </button>
            </div>
        </nav>
    )
}

export default Topbar