import { deleteLanguage, installLanguage, uninstallLanguage } from '../services/api';

function LanguageItem({ lang, refresh }) {
    function formatDate(isoDate) {
        return new Date(isoDate).toLocaleString();
    }

    return (
        <div
            className='border p-2 rounded w-[500px] mb-4'
        >
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
                <button onClick={() => { deleteLanguage(lang.id).then(refresh); }} className='px-4 py-2 rounded bg-red-400 text-white hover:bg-red-700'>Delete</button>{' '}
                {lang.installed ? (
                    <button onClick={() => { uninstallLanguage(lang.id).then(refresh); }} className='px-4 py-2 rounded bg-red-400 text-white hover:bg-red-700'>Uninstall</button>
                ) : (
                    <button onClick={() => { installLanguage(lang.id).then(refresh); }} className='px-4 py-2 rounded bg-green-400 text-white hover:bg-green-700'>Install</button>
                )}
            </div>
        </div>
    )
}

export default LanguageItem