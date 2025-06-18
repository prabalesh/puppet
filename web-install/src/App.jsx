import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import { useState } from 'react';
import LanguageForm from './components/LanguageForm';
import LanguageList from './components/LanguageList';
import CodeExecution from './pages/CodeExecution';

export default function App() {
  const [refreshKey, setRefreshKey] = useState(0);

  return (
    <Router>
      <div style={{ padding: '1rem' }}>
        <h1>Puppet</h1>
        <nav style={{ marginBottom: '1rem' }}>
          <Link to="/" style={{ marginRight: 10 }}>Language Manager</Link>
          <Link to="/executions">Run Code</Link>
        </nav>

        <Routes>
          <Route
            path="/"
            element={
              <>
                <LanguageForm onSuccess={() => setRefreshKey(k => k + 1)} />
                <LanguageList key={refreshKey} />
              </>
            }
          />
          <Route path="/executions" element={<CodeExecution />} />
        </Routes>
      </div>
    </Router>
  );
}
