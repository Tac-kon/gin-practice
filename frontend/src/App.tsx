// App.tsx
import './App.css';
import {BrowserRouter, Routes, Route} from 'react-router-dom';
import Users from './pages/Users';
import Login from './pages/Login';
import Register from './pages/Register';
import {RedirectToUsers} from './components/RedirectToUsers';
import Links from './pages/Links';

function App() {
  return (
    // <div className="App">
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<RedirectToUsers />} />
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} /> 
        <Route path={'/users'} element={<Users />} />
        <Route path={'/users/:id/links'} element={<Links />} />
      </Routes>
    </BrowserRouter>
    // </div>
  );
}
export default App;