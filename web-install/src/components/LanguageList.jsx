import { useEffect, useState } from 'react';
import {
  getLanguages,
  deleteLanguage,
  installLanguage,
  uninstallLanguage,
} from '../services/api';

export default function LanguageList() {
  const [languages, setLanguages] = useState([]);

  async function refresh() {
    const data = await getLanguages();
    setLanguages(data);
  }

  useEffect(() => {
    refresh();
  }, []);

  function formatDate(isoDate) {
    return new Date(isoDate).toLocaleString();
  }

  return (
    <div>
      {languages != null && languages.length > 0 ? (
        <>
          {languages.map((lang) => (
            <div
              key={lang.id}
              style={{
                border: '1px solid #ccc',
                padding: '0.75rem',
                marginBottom: '0.75rem',
                borderRadius: '4px',
              }}
            >
              <h4>{lang.name} ({lang.version})</h4>
              <p><strong>Status:</strong> {lang.installed ? '✅ Installed' : '❌ Not installed'}</p>
              <p><strong>Docker Image:</strong> {lang.image_name}</p>
              <p><strong>Source File Name:</strong> {lang.file_name}</p>
              <p><strong>Compile Command:</strong> {lang.compile_command || '—'}</p>
              <p><strong>Run Command:</strong> {lang.run_command}</p>
              <p><strong>Created At:</strong> {formatDate(lang.created_at)}</p>

              <div style={{ marginTop: '0.5rem' }}>
                <button onClick={() => { deleteLanguage(lang.id).then(refresh); }}>Delete</button>{' '}
                {lang.installed ? (
                  <button onClick={() => { uninstallLanguage(lang.id).then(refresh); }}>Uninstall</button>
                ) : (
                  <button onClick={() => { installLanguage(lang.id).then(refresh); }}>Install</button>
                )}
              </div>
            </div>
          ))}
        </>
      ) : (
        <p>No languages to show</p>
      )}
    </div>
  );
}
