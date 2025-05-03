import React, { useState } from "react";
import axios from "axios";
import toast from "react-hot-toast";
import { LoggedUser, MountedPartition, User } from "../../App";

interface LogInProps {
    onLogin: (user: User, partition: MountedPartition) => void; // Callback to set user information in App.tsx
    defaultIdPartition?: string; // Optional default partition ID
    onClose: () => void; // Callback to close the overlay
}

export default function LogIn({ onLogin, defaultIdPartition = "", onClose }: LogInProps) {
    const [idPartition, setIdPartition] = useState<string>(defaultIdPartition);
    const [username, setUsername] = useState<string>("");
    const [password, setPassword] = useState<string>("");
    const [loading, setLoading] = useState<boolean>(false);

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();

        if (!idPartition || !username || !password) {
            toast.error("All fields are required!");
            return;
        }

        setLoading(true);

        try {
            const parseResponse = await axios.post("http://localhost:8080/parse", { code: "\nmounted\nlogin -id=" + idPartition + " -user=" + username + " -pass=" + password + "\nmounted" }, {
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


            // Step 2: Proceed to login with the provided credentials
            const loginResponse = await axios.get<LoggedUser>("http://localhost:8080/login", {
                params: {
                    idPartition: idPartition,
                    username: username,
                    password: password,
                },
            });

            const { User, MountedPartition } = loginResponse.data;

            // Handle success
            onLogin(User, MountedPartition); // Set user information in App.tsx
            toast.success("Login successful!");
            onClose(); // Close the overlay
        } catch (error) {
            // Handle errors from either parse or login requests
            toast.error("Login failed. Check the credentials or server status.");
            console.error("Login Error:", error);
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="fixed inset-0 bg-gray-900 bg-opacity-75 flex items-center justify-center z-50">
            <div className="bg-white text-black p-6 rounded-lg shadow-lg w-full max-w-sm">
                <h2 className="text-xl font-bold mb-4">Log in to Partition</h2>
                <form onSubmit={handleSubmit} className="space-y-4">
                    <div>
                        <label htmlFor="idPartition" className="block text-sm font-medium">
                            Partition ID
                        </label>
                        <input
                            type="text"
                            id="idPartition"
                            value={idPartition}
                            onChange={(e) => setIdPartition(e.target.value)}
                            className="w-full p-2 border rounded mt-1"
                        />
                    </div>
                    <div>
                        <label htmlFor="username" className="block text-sm font-medium">
                            Username
                        </label>
                        <input
                            type="text"
                            id="username"
                            value={username}
                            onChange={(e) => setUsername(e.target.value)}
                            className="w-full p-2 border rounded mt-1"
                        />
                    </div>
                    <div>
                        <label htmlFor="password" className="block text-sm font-medium">
                            Password
                        </label>
                        <input
                            type="password"
                            id="password"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            className="w-full p-2 border rounded mt-1"
                        />
                    </div>
                    <div className="flex items-center justify-between">
                        <button
                            type="submit"
                            disabled={loading}
                            className="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700 disabled:opacity-50"
                        >
                            {loading ? "Logging in..." : "Log In"}
                        </button>
                        <button
                            type="button"
                            onClick={onClose}
                            className="px-4 py-2 bg-gray-600 text-white rounded hover:bg-gray-700"
                        >
                            Cancel
                        </button>
                    </div>
                </form>
            </div>
        </div>
    );
}