import React, { useState } from 'react';
import { FaFolder, FaFile } from 'react-icons/fa';

export interface FileItem {
    name: string;
    isDirectory: boolean;
    size?: number;
    content?: string;
    children?: FileItem[];
}

interface FileExplorerProps {
    files: FileItem[];
    onOpenFile?: (file: FileItem) => void;
}

const FileExplorer: React.FC<FileExplorerProps> = ({ files, onOpenFile }) => {
    const [currentFiles, setCurrentFiles] = useState<FileItem[]>(files);
    const [pathStack, setPathStack] = useState<FileItem[]>([]);
    const [selectedItem, setSelectedItem] = useState<FileItem | null>(null);

    const handleItemClick = (item: FileItem) => {
        setSelectedItem(item);
        if (item.isDirectory) {
            setPathStack(prev => [...prev, item]);
            setCurrentFiles(item.children || []);
        } else {
            if (onOpenFile) {
                onOpenFile(item);
            }
        }
    };

    const goUp = () => {
        if (pathStack.length > 0) {
            const newStack = [...pathStack];
            newStack.pop();
            setPathStack(newStack);
            setSelectedItem(null);
            if (newStack.length === 0) {
                setCurrentFiles(files);
            } else {
                const parent = newStack[newStack.length - 1];
                setCurrentFiles(parent.children || []);
            }
        }
    };

    return (
        <div style={{ display: 'flex', height: '100%' }}>
            {/* Sidebar with file/folder icons */}
            <div style={{ flex: 1, borderRight: '1px solid #ccc', padding: '1rem', display: 'flex', flexDirection: 'column' }}>
                <button onClick={goUp} disabled={pathStack.length === 0}>Up</button>
                <div style={{
                    display: 'grid',
                    gridTemplateColumns: 'repeat(auto-fill, minmax(100px, 1fr))',
                    gap: '1rem',
                    marginTop: '1rem',
                    overflowY: 'auto'
                }}>
                    {currentFiles.map(item => (
                        <div
                            key={item.name}
                            onClick={() => handleItemClick(item)}
                            style={{ cursor: 'pointer', textAlign: 'center' }}
                        >
                            {item.isDirectory ? <FaFolder size={48} /> : <FaFile size={48} />}
                            <div
                                title={item.name}
                                style={{ whiteSpace: 'nowrap', overflow: 'hidden', textOverflow: 'ellipsis' }}
                            >
                                {item.name}
                            </div>
                        </div>
                    ))}
                </div>
            </div>

            {/* Details panel */}
            <div style={{ width: '300px', padding: '1rem' }}>
                {selectedItem ? (
                    <div>
                        <h3>{selectedItem.name}</h3>
                        <p>Type: {selectedItem.isDirectory ? 'Directory' : 'File'}</p>
                        {!selectedItem.isDirectory && <p>Size: {selectedItem.size} bytes</p>}
                        {!selectedItem.isDirectory && (
                            <div style={{
                                marginTop: '1rem',
                                maxHeight: 'calc(100% - 4rem)',
                                overflow: 'auto',
                                background: '#f9f9f9',
                                padding: '0.5rem',
                                border: '1px solid #ddd'
                            }}>
                                <pre>{selectedItem.content}</pre>
                            </div>
                        )}
                        {selectedItem.isDirectory && <p>Items: {(selectedItem.children || []).length}</p>}
                    </div>
                ) : (
                    <p>Select a file or folder</p>
                )}
            </div>
        </div>
    );
};

export default FileExplorer;

/**
 * Example usage:
 *
 * const files: FileItem[] = [
 *   {
 *     name: 'Documents',
 *     isDirectory: true,
 *     children: [
 *       { name: 'Resume.pdf', isDirectory: false, size: 102400, content: '...' },
 *       { name: 'Projects', isDirectory: true, children: [] }
 *     ]
 *   },
 *   { name: 'readme.txt', isDirectory: false, size: 1024, content: 'Hello world' },
 * ];
 *
 * <FileExplorer files={files} onOpenFile={file => console.log('Opening', file.name)} />
 */
