import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

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
			navigate("/chat");
		}
	};


	return (
		<div>
			<p>login: </p>
			<form onSubmit={handleFormSubmit}>
				<input 
					type="text" 
					value={loginName}
					onChange={handleInputChange}
				/>
			</form>
		</div>
	);
};

export default Login;
