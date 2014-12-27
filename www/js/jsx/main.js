/** @jsx React.DOM */
var Content = React.createClass({
  getInitialState: function() {
    return {data: []}
  },
  componentDidMount: function() {
  },
  handleQuerySubmit: function(query) {
      $.ajax({
      url: this.props.url,
      type: "GET",
      dataType: 'json',
      data: {q: query},
      success: function(data) {
        this.setState({data: data})
      }.bind(this)
    });
  },
  render: function() {
    return (
      <div className="content">
        <h1>Search</h1>
        <SearchForm onQuerySubmit={this.handleQuerySubmit}/>
        <ItemList data={this.state.data} />
      </div>
    );
  }
});

var SearchForm = React.createClass({
  handleSubmit: function(e) {
    e.preventDefault();
    var query = this.refs.query.getDOMNode().value.trim();
    if (!query) return;
    this.props.onQuerySubmit(query);
  },
  render: function() {
    return (
      <form className="searchForm" method="GET" action="/search" onSubmit={this.handleSubmit}>
        <input id="searchWord" placeholder="niku" ref="query"></input>
        <button id="searchButton"><span>search</span></button>
      </form>
    );
  }
});

var ItemList = React.createClass({
  render: function() {
    if (!this.props.data) return;
    var itemNodes = this.props.data.map(function(item) {
      return (
        <Item author={item.Author} title={item.Title} beginning={item.Beginning}></Item>
      );
    });
    return (
      <div className="itemList">
       {itemNodes}
     </div>
    );
  }
});

var Item = React.createClass({
  render: function() {
    return (
      <ul className="item">
        <li className="author">
          {this.props.author}
        </li>
        <li className="title">
          {this.props.title}
        </li>
        <li className="beginning">
          {this.props.beginning}
        </li>
      </ul>
    );
  }
});

React.render(
  <Content url="/api/v0/search"/>,
  document.getElementById('content')
);