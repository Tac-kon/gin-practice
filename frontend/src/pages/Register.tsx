// pages/Register.tsx
import React, {Component, SyntheticEvent} from "react";
import axios from "axios"
import {User} from '../models/user'
import {Navigate} from 'react-router-dom' 

class Register extends Component {
	firstName = ''
	lastName = ''
	email = ''
	password = ''
	passwordConfirm = ''

	registerUrl = 'register'
	state = {
		redirect: false
	}

	submit = async (e: SyntheticEvent) => {
		e.preventDefault()
		
		const user = new User(
			this.firstName,
			this.lastName,
			this.email,
			this.password,
			this.passwordConfirm
		)
		
		await axios.post(this.registerUrl, user)

		this.setState({
			redirect: true
		})
	}
	
	render() {
		if (this.state.redirect) {
			return <Navigate to='/login' />
		}
		return (
			<main className="form-signin">
				<form onSubmit={this.submit}>
					<h1 className="h3 mb-3 fw-normal">Please register</h1>
					
					<div className="form-floating">
						<input className="form-control" placeholder="First Name"
							onChange={e => {
								this.submit(e);
								this.firstName = e.target.value;
							}}
						/>
						<label>First Name</label>
					</div>
					
					<div className="form-floating">
						<input className="form-control" placeholder="Last Name"
							onChange={e => {
								this.submit(e);
								this.lastName = e.target.value;
							}}
						/>
						<label>Last Name</label>
					</div>
					
					<div className="form-floating">
						<input type="email" className="form-control" placeholder="name@example.com"
							onChange={e => {
								this.submit(e);
								this.email = e.target.value;
							}}
						/>
						<label>Email address</label>
					</div>
					
					<div className="form-floating">
						<input type="password" className="form-control" placeholder="Password"
							onChange={e => {
								this.submit(e);
								this.password = e.target.value;
							}}
						/>
						<label>Password</label>
					</div>
					
					<div className="form-floating">
						<input type="password" className="form-control" placeholder="Password Confirm"
							onChange={e => {
								this.submit(e);
								this.passwordConfirm = e.target.value;
							}}
						/>
						<label>Password Confirm</label>
					</div>
					
					<button className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
				</form>
			</main>
		)
	}
}
export default Register