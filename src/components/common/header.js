import React, { PropTypes } from 'react'

const Header = React.createClass({
  render () {
    return (
      <nav className="navbar navbar-default">
        <div className="container-fluid">
          <a href="/" className="navbar-brand">
            <img src="images/logo.png" height="50px"/>
          </a>
          <ul className="nav navbar-nav">
            <li><a href="/">Home</a></li>
          </ul>
        </div>
      </nav>
    )
  }
})

export default Header
