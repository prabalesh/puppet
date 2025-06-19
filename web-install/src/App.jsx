import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import { useState } from "react";
import LanguageForm from "./components/LanguageForm";
import LanguageList from "./components/LanguageList";
import CodeExecution from "./pages/CodeExecution";

export default function App() {
    const [refreshKey, setRefreshKey] = useState(0);

    return (
        <Router>
            <div>
                <header className="mb-5 h-[10vh] bg-slate-600 text-white text-2xl flex justify-around items-center">
                    <div>
                        <h1>Puppet</h1>
                    </div>
                    <nav className="flex gap-4 text-xl">
                        <Link to="/" style={{ marginRight: 10 }}>
                            Language Manager
                        </Link>
                        <Link to="/executions">Run Code</Link>
                    </nav>
                </header>

                <main className="container mx-auto">
                    <Routes>
                        <Route
                            path="/"
                            element={
                                <>
                                    <LanguageForm
                                        onSuccess={() =>
                                            setRefreshKey((k) => k + 1)
                                        }
                                    />
                                    <LanguageList key={refreshKey} />
                                </>
                            }
                        />
                        <Route path="/executions" element={<CodeExecution />} />
                    </Routes>
                </main>
            </div>
        </Router>
    );
}
