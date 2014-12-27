/** @jsx React.DOM */
var Content = React.createClass({displayName: "Content",
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
      React.createElement("div", {className: "content"}, 
        React.createElement("h1", null, "Search"), 
        React.createElement(SearchForm, {onQuerySubmit: this.handleQuerySubmit}), 
        React.createElement(ItemList, {data: this.state.data})
      )
    );
  }
});

var SearchForm = React.createClass({displayName: "SearchForm",
  handleSubmit: function(e) {
    e.preventDefault();
    var query = this.refs.query.getDOMNode().value.trim();
    if (!query) return;
    this.props.onQuerySubmit(query);
  },
  render: function() {
    return (
      React.createElement("form", {className: "searchForm", method: "GET", action: "/search", onSubmit: this.handleSubmit}, 
        React.createElement("input", {id: "searchWord", placeholder: "niku", ref: "query"}), 
        React.createElement("button", {id: "searchButton"}, React.createElement("span", null, "search"))
      )
    );
  }
});

var ItemList = React.createClass({displayName: "ItemList",
  render: function() {
    if (!this.props.data) return;
    var itemNodes = this.props.data.map(function(item) {
      return (
        React.createElement(Item, {author: item.Author, title: item.Title, beginning: item.Beginning})
      );
    });
    return (
      React.createElement("div", {className: "itemList"}, 
       itemNodes
     )
    );
  }
});

var Item = React.createClass({displayName: "Item",
  render: function() {
    return (
      React.createElement("ul", {className: "item"}, 
        React.createElement("li", {className: "author"}, 
          this.props.author
        ), 
        React.createElement("li", {className: "title"}, 
          this.props.title
        ), 
        React.createElement("li", {className: "beginning"}, 
          this.props.beginning
        )
      )
    );
  }
});

React.render(
  React.createElement(Content, {url: "/api/v0/search"}),
  document.getElementById('content')
);