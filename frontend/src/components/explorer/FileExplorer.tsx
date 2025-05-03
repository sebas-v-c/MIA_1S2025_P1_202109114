import { useState } from 'react';

export type FileType = 'File' | 'Directory' | 'Disc' | 'Partition';

export type File = {
    name: string;
    size: number;
    type: FileType;
    createdTime: string;
    modifiedTime: string;
    content?: string;
    // Disc-specific fields
    mbrSize?: number;
    signature?: number;
    fit?: string;
    // Partition-specific fields
    partitions?: File[];
    status?: string;
    start?: number;
    correlative?: number;
    id?: string;
    partType?: string;
};

export interface FileExplorerProps {
    files: File[];
    onFolderOpen?: (item: File) => void;
    onFileOpen?: (item: File) => void;
    onFolderDoubleClick?: (item: File) => void;
    onFileDoubleClick?: (item: File) => void;
    onNavigateUp?: () => void;
}


const getIcon = (type: FileType) => {
    switch (type) {
        case 'Directory': return '📁';
        case 'Disc':      return '💽';
        case 'Partition': return '📀';
        default:          return '📄';
    }
};

export default function FileExplorer({
                                         files,
                                         onFolderOpen,
                                         onFileOpen,
                                         onFolderDoubleClick,
                                         onFileDoubleClick,
                                         onNavigateUp,
                                     }: FileExplorerProps) {
    const [selected, setSelected] = useState<File | null>(null);

    const handleClick = (f: File) => {
        setSelected(f);
        if (f.type === 'Directory') onFolderOpen?.(f);
        //else onFileOpen?.(f);
    };

    const handleDoubleClick = (f: File) => {
        if (f.type === 'Directory') onFolderDoubleClick?.(f);
        else if (f.type === 'Disc') onFileDoubleClick?.(f);
        else if (f.type === 'Partition') onFileDoubleClick?.(f);
    };


    return (
        <div className="flex h-full w-full font-sans">
            {/* Sidebar */}
            <div className="flex-1 border-r border-gray-800 bg-gray-900 overflow-y-auto p-4">
                <button
                    onClick={onNavigateUp}
                    disabled={!onNavigateUp}
                    className="mb-4 px-4 py-2 rounded bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50"
                >
                    ⬆️ Up
                </button>
                <ul className="space-y-2">
                    {files.map(f => (
                        <li
                            key={`${f.name}-${f.createdTime}`}
                            onClick={() => handleClick(f)}
                            onDoubleClick={() => handleDoubleClick(f)}
                            className={`flex items-center p-2 rounded cursor-pointer hover:bg-blue-800 ${
                                selected === f ? 'bg-blue-900 text-white' : 'text-gray-100'
                            }`}
                        >
                            <span className="mr-2">{getIcon(f.type)}</span>
                            {f.name}
                        </li>
                    ))}
                </ul>
            </div>

            {/* Info panel */}
            <div className="flex-none w-1/3 max-w-xs p-4 bg-gray-800 text-gray-100 overflow-y-auto">
                {selected ? (
                    <div className="space-y-4">
                        <h3 className="text-xl font-semibold">{selected.name}</h3>
                        <ul className="space-y-2">
                            <li><strong>Type:</strong> {selected.type}</li>
                            {selected.type === 'Disc' && (
                                <>
                                    <li><strong>MBR Size:</strong> {selected.mbrSize} bytes</li>
                                    <li><strong>Signature:</strong> {selected.signature}</li>
                                    <li><strong>Fit:</strong> {selected.fit}</li>
                                    <li><strong>Partitions:</strong> {selected.partitions?.length}</li>
                                </>
                            )}
                            {selected.type === 'Partition' && (
                                <>
                                    <li><strong>Status:</strong> {selected.status}</li>
                                    <li><strong>Start:</strong> {selected.start}</li>
                                    <li><strong>Correlative:</strong> {selected.correlative}</li>
                                    <li><strong>ID:</strong> {selected.id}</li>
                                    <li><strong>Partition Type:</strong> {selected.partType}</li>
                                    <li><strong>Size:</strong> {selected.size}</li>
                                </>
                            )}
                            <li><strong>Created:</strong> {selected.createdTime}</li>
                            <li><strong>Modified:</strong> {selected.modifiedTime}</li>
                        </ul>
                        {selected.content !== undefined && (
                            <div>
                                <strong>Content:</strong>
                                <textarea
                                    readOnly
                                    className="w-full h-32 mt-1 p-2 bg-gray-700 text-white rounded resize-none"
                                    value={selected.content}
                                />
                            </div>
                        )}
                    </div>
                ) : (
                    <div className="text-gray-500">Select an item to see details</div>
                )}
            </div>
        </div>
    );
}
