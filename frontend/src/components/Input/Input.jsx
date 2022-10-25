import React from "react";
import DatePicker from "react-datepicker";
// import NumberFormat from "react-number-format";

import "react-datepicker/dist/react-datepicker.css";

export default class Input extends React.Component {
	constructor(props) {
		super(props);

		const d = new Date();

		this.state = {
			date: d,
		}
	}

	loadData = async () => {
		// todo
	}

	render = () => {
		return (<>
			<h3>Navigation Header Here</h3>
			<h1>Calorie Input Here</h1>
			<div>
				Date: <DatePicker selected={this.state.date} onChange={(d)=>this.setState({date:d})} /><br/>
				Item Name: <input list="types" /><br/>
				<datalist id="types"></datalist>
				Calories: <input type="text"/><br/>
				Servings: <input type="text"/><br/>
			</div>
		</>);
	}
}
