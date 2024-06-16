import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Login from "./Login";
import ChatWindow from "./ChatWindow";
import './App.css';

function App() {
	return (
		<Router>
			<div className="App">
				<Routes>
					<Route path="/" element={<Login />} />
					<Route path="/chat" element={<ChatWindow />} />
				</Routes>
			</div>
		</Router>
	);
}

export default App;
