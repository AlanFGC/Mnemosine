import { BrowserRouter, Routes, Route } from 'react-router-dom';
import EditPage from './pages/EditPage/EditPage';
import Navbar from './components/NavBar/NavBar';
import './App.css';

function App() {
  return (
    <div className="app-container">
      <Navbar />
      <BrowserRouter>
        <Routes>
          <Route path="/editor/" element={<EditPage />} />
          <Route path="/editor/:id" element={<EditPage />} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
