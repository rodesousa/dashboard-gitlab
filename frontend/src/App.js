import React, { Component, Fragment } from 'react';
import logo from './logo.svg';
import './App.css';
import './msg.css';


class App extends Component {

  render(){
    var fuck =
      <div className="success-msg">
        <i className="fa fa-check"></i>
        <a href={this.props.merge_request.web_url}>{this.props.merge_request.title}</a>
      </div>

    return (
      <Fragment>
        {fuck}
      </Fragment>
    )
  }

}

export default App;
