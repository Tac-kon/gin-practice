// App.tsx
import './App.css';
import {BrowserRouter, Route} from 'react-router-dom';
import Users from './pages/Users';
import Login from './pages/Login';
import Register from './pages/Register';

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Route path="/" element={<Users />} />
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} /> 
      </BrowserRouter>
    </div>
  );
}
export default App;