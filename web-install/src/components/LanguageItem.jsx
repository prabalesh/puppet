import { useState } from 'react';
import {
    deleteLanguage,
    installLanguage,
    uninstallLanguage,
    getJobStatus
} from '../services/api';

function LanguageItem({ lang, refresh }) {
    const [isProcessing, setIsProcessing] = useState(false);
    const [processingAction, setProcessingAction] = useState("");

    function formatDate(isoDate) {
        return new Date(isoDate).toLocaleString();
    }

    async function pollJobUntilDone(jobId) {
        let attempts = 0;
        const maxAttempts = 30;
        const interval = 3000;

        return new Promise((resolve, reject) => {
            const poll = async () => {
                try {
                    const status = await getJobStatus(jobId);
                    if (status.status === 'done') {
                        resolve();
                    } else if (attempts++ >= maxAttempts) {
                        reject(new Error("Polling timed out"));
                    } else {
                        setTimeout(poll, interval);
                    }
                } catch (err) {
                    reject(err);
                }
            };
            poll();
        });
    }

    async function handleJob(actionFunc, actionType) {
        try {
            setIsProcessing(true);
            setProcessingAction(actionType);
            const { job_id } = await actionFunc();
            if(actionType == "delete" && job_id == -1) {
                refresh()
                return
            }
            await pollJobUntilDone(job_id);
            refresh();
        } catch (err) {
            console.error(`Failed to ${actionType}:`, err);
        } finally {
            setIsProcessing(false);
            setProcessingAction("");
        }
    }

    return (
        <div className='border p-2 rounded mb-4'>
            <div className='mb-4 text-xl font-bold text-center'>
                <h4>{lang.name} ({lang.version})</h4>
            </div>
            <div className='gap-2 mb-4'>
                <p><strong>Status:</strong> {lang.installed ? '✅ Installed' : '❌ Not installed'}</p>
                <p><strong>Docker Image:</strong> {lang.image_name}</p>
                <p><strong>Source File Name:</strong> {lang.file_name}</p>
                <p><strong>Compile Command:</strong> {lang.compile_command || '—'}</p>
                <p><strong>Run Command:</strong> {lang.run_command}</p>
                <p><strong>Created At:</strong> {formatDate(lang.created_at)}</p>
            </div>

            <div style={{ marginTop: '0.5rem' }}>
                <button
                    disabled={isProcessing}
                    onClick={() => handleJob(() => deleteLanguage(lang.id), "delete")}
                    className='px-4 py-2 rounded bg-red-400 text-white hover:bg-red-700 disabled:opacity-50 disabled:cursor-not-allowed'
                >
                    {isProcessing && processingAction === "delete" ? "Deleting..." : "Delete"}
                </button>{' '}

                {lang.installed ? (
                    <button
                        disabled={isProcessing}
                        onClick={() => handleJob(() => uninstallLanguage(lang.id), "uninstall")}
                        className='px-4 py-2 rounded bg-red-400 text-white hover:bg-red-700 disabled:opacity-50 disabled:cursor-not-allowed'
                    >
                        {isProcessing && processingAction === "uninstall" ? "Uninstalling..." : "Uninstall"}
                    </button>
                ) : (
                    <button
                        disabled={isProcessing}
                        onClick={() => handleJob(() => installLanguage(lang.id), "install")}
                        className='px-4 py-2 rounded bg-green-400 text-white hover:bg-green-700 disabled:opacity-50 disabled:cursor-not-allowed'
                    >
                        {isProcessing && processingAction === "install" ? "Installing..." : "Install"}
                    </button>
                )}
            </div>
        </div>
    );
}

export default LanguageItem;
