import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import "./Login.css";

const Login = () => {

	const [loginName, setLoginName] = useState("");
	// useNavigateフックを使って、navigate関数を取得します。これにより、プログラム的に他のルートにナビゲートできる
	const navigate = useNavigate();

	const handleInputChange = (e) => {
		setLoginName(e.target.value);
	};

	const handleFormSubmit = (e) => {
		// ブラウザのデフォルトのフォーム送信動作を防止
		e.preventDefault();
		// ChatWindowコンポーネントに移動
		if (loginName.trim()) {
			// loginNameをURLパラメーターとしてChatWindowに渡す
			navigate(`/chat?name=${loginName}`);
		}
	};


	return (
		<div className="login-container">
			<div className="login-form">
				<h1>Login</h1>
				<form onSubmit={handleFormSubmit}>
					<input
						type="text"
						value={loginName}
						onChange={handleInputChange}
						placeholder="Enter your name"
					/>
					<button type="submit">Login</button>
				</form>
			</div>
		</div>
	);
};

export default Login;
