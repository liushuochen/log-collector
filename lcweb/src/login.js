import React from "react";
import './Login.css';

class Login extends React.Component {
	render() {
		return (
			<div className="App">
				<header className="App-header">
					<form className="App-form">
						<div>
							<input placeholder="username" type="text" name="name" />
						</div>
						<div>
							<input placeholder="password" type="password" name="pwd" />
						</div>
						<div>
							<input type="submit" value="sign in" />
						</div>
						<div>
							<input type="button" value={'register'} onClick={(e) => this.registerFunc(e.target.value)} />
						</div>
					</form>
				</header>
			</div>
		)
	}

	registerFunc = (name) => {
		console.log(name);
	}
}

export default Login;