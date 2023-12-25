import { BrowserRouter, Routes, Route } from 'react-router-dom';
import EditPage from './pages/EditPage/EditPage';

function App() {
  return (
    <div className="App">
      <h1>Welcome to Mnesomine</h1>
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
