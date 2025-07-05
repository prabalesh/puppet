import { useState } from 'react';
import {
    deleteLanguage,
    installLanguage,
    uninstallLanguage,
    getJobStatus
} from '../services/api';
import { Download, Trash2, XCircle } from 'lucide-react';
import Spinner from './Spinner';

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
            if (actionType === "delete" && job_id === -1) {
                refresh();
                return;
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
                <p className='text-xs font-light'>{formatDate(lang.created_at)}</p>
            </div>
            <div className='gap-2 mb-4 text-center'>
                <p><strong>Status:</strong> {lang.installed ? '✅ Installed' : '❌ Not installed'}</p>
                <p><strong>Docker Image:</strong> {lang.image_name}</p>
                <p><strong>Source File Name:</strong> {lang.file_name}</p>
                <p><strong>Compile Command:</strong> {lang.compile_command || '—'}</p>
                <p><strong>Run Command:</strong> {lang.run_command}</p>
            </div>

            <div className='text-center flex justify-center gap-4 mt-2'>
                <button
                    disabled={isProcessing}
                    onClick={() => handleJob(() => deleteLanguage(lang.id), "delete")}
                    className='flex items-center gap-2 px-4 py-2 rounded border border-red-400 text-red-400 hover:bg-red-400 hover:text-white disabled:opacity-50 disabled:cursor-not-allowed'
                    title='Delete Language'
                >
                    {isProcessing && processingAction === "delete" ? <Spinner /> : <XCircle size={18} />}
                </button>

                {lang.installed ? (
                    <button
                        disabled={isProcessing}
                        onClick={() => handleJob(() => uninstallLanguage(lang.id), "uninstall")}
                        className='flex items-center gap-2 px-4 py-2 rounded border border-red-400 text-red-400 hover:bg-red-400 hover:text-white disabled:opacity-50 disabled:cursor-not-allowed'
                        title='Uninstall Language'
                    >
                        {isProcessing && processingAction === "uninstall" ? <Spinner /> : <Trash2 size={18} />}
                    </button>
                ) : (
                    <button
                        disabled={isProcessing}
                        onClick={() => handleJob(() => installLanguage(lang.id), "install")}
                        className='flex items-center gap-2 px-4 py-2 rounded border border-green-400 text-green-400 hover:bg-green-400 hover:text-white disabled:opacity-50 disabled:cursor-not-allowed'
                        title='Install Language'
                    >
                        {isProcessing && processingAction === "install" ? <Spinner /> : <Download size={18} />}
                    </button>
                )}
            </div>
        </div>
    );
}

export default LanguageItem;
