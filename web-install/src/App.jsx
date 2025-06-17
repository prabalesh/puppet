import { useState } from 'react';
import LanguageForm from './components/LanguageForm';
import LanguageList from './components/LanguageList';

export default function App() {
  const [refreshKey, setRefreshKey] = useState(0);
  return (
    <div style={{ padding: '1rem' }}>
      <h1>Puppet | Language Manager</h1>
      <LanguageForm onSuccess={() => setRefreshKey(k => k + 1)} />
      <LanguageList key={refreshKey} />
    </div>
  );
}