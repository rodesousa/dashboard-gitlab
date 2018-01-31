import React, { Component, Fragment } from 'react';
import logo from './logo.svg';
import App from './App';
import './Accordeon.css';

class Accordeon extends Component {

  constructor(props){
    super(props);
    this.state = {
      repositories: [{}],
    }
  }

  majState(){

    fetch(
      'http://localhost:8080/merge_request/all')
      .then((response) => response.json())
      .then((responseJson) => {
        this.setState({
          repositories: responseJson.merge_requests,
        } )
      })
  }

  componentDidMount(){
    this.majState()
  }

  click = (i) => {
    const newAccordion = this.state.repositories.slice();
    const index = newAccordion.indexOf(i)

    newAccordion[index].open = !newAccordion[index].open;
    this.setState({repositories: newAccordion});
  }

  reloadApiRepo = () => {
    fetch('http://localhost:8080/merge_request/maj')
    .then((response) => this.majState())
  }

  render () {


    function bodyApp(repository) {
      if('MergeRequests' in repository){
        return repository.MergeRequests.map((mr) => {
          return <App merge_request={mr} />
        })
      }
    }

    function MergeRequestsCount(repository){
      if('MergeRequests' in repository){
          return repository.MergeRequests.length
      }
    }

    const sections = this.state.repositories.map((i) => (
      <div key={this.state.repositories.indexOf(i)}>
      <div
      className="title"
      onClick={this.click.bind(null, i)}
      >
      <div className="arrow-wrapper">
      <i className={i.open
        ? "fa fa-angle-down fa-rotate-180"
        : "fa fa-angle-down"}
        ></i>
        </div>
        <span className="title-text">
          {i.name} ({MergeRequestsCount(i)})
        </span>
        </div>
        <div className={i.open
          ? "content content-open"
          : "content"}
          >
          <div className={i.open
            ? "content-text content-text-open"
            : "content-text"}>
            {bodyApp(i)}
            </div>
            </div>
            </div>
          ));

          return (
            <div className="accordion">
            {sections}
            <button onClick={this.reloadApiRepo}>
              Reload !
            </button>
            </div>
          );
        }
      }

export default Accordeon;
