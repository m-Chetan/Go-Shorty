import {Route, BrowserRouter as Router, Routes} from "react-router-dom"
import Home from "./screens/Home";
import Header from "./components/Header";

function App() {
  return (
    <Router>
      <Header />
      <Routes>
        {/* <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} /> */}
        <Route path="/" element={<Home />} />
      </Routes>
    </Router>
    
  );
}

export default App;
