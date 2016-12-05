import $ from 'jquery';
import jQuery from 'jquery';
import ReactDOM from 'react-dom';
import React, { PropTypes } from 'react'
import Home from './components/homePage';
import Header from './components/common/header';

const App = React.createClass({
  render () {
    let Child;
    switch (this.props.route) {
      default: Child = Home;

    }
    return (
      <div>
        <Header/>
        <Child/>
      </div>
    )
  }
})

function render () {
  let route = window.location.hash.substr(1);
  ReactDOM.render(<App route={route} />, document.getElementById('app'));
}
window.addEventListener('hashchange', render);
render();
