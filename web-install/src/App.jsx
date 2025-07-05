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
                <header className="mb-5 h-[10vh] bg-white border-b-1 font-bold text-purple-600 text-2xl flex justify-around items-center">
                    <div>
                        <Link to={"/"}>
                            <img src="./puppet.png" className="h-36 w-36" alt="Puppet" />
                        </Link>
                    </div>
                    <nav className="flex gap-6 text-xl">
                        <Link to="/" className="hover:text-purple-800">
                            Language Manager
                        </Link>
                        <Link to="/executions" className="hover:text-purple-800">Run Code</Link>
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
