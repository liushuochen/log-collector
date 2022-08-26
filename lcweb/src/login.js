import React from "react";
import './Login.css';

class Login extends React.Component {
	render() {
		return (
			<div className="App">
				<header className="App-header">
					<form className="App-form">
						<div>
							<input placeholder="user name" type="text" name="name" />
						</div>
						<div>
							<input placeholder="password" type="password" name="pwd" />
						</div>
						<input type="submit" value="Sign in" />
					</form>
				</header>
			</div>
		)
	}
}

export default Login;