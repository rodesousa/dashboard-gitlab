import React, { Component, Fragment } from 'react';
import logo from './logo.svg';

class Token extends Component {

  constructor(props){
    super(props);
    this.state = {
      value: ''
     }
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(event) {
    this.setState({value: event.target.value});
  }

  handleSubmit(event) {
    fetch('http://localhost:8080/api/token?value='+this.state.value,{
      method: "POST"
    })
    event.preventDefault();
  }

  render(){

    var form = <form onSubmit={this.handleSubmit}>
      <label>
        Token:
        <input type="password" value={this.state.value} onChange={this.handleChange} />
      </label>
    </form>

    return (
      <Fragment>
        {form}
      </Fragment>
    )
  }

}

export default Token;
