import React, { Component } from "react";
import axios from "axios"; // It makes the API calls to the backend
import { Card, Header, Form, Input, Icon, Button } from "semantic-ui-react"; // Import Button
import "./ToDoList.css"; // Import the custom CSS file
let endpoint = "http://localhost:9000";

class ToDoList extends Component {
  constructor(props) {
    super(props);

    this.state = {
      task: "",
      items: [],
    };
  }

  componentDidMount() {
    this.getTask();
  }

  onChange = (event) => {
    this.setState({
      [event.target.name]: event.target.value,
    });
  };

  onSubmit = (event) => {
    event.preventDefault(); // Prevent default form submission behavior
    let { task } = this.state;

    if (task) {
      axios
        .post(
          endpoint + "/api/task",
          { task: task }, // Ensure the task is correctly sent
          {
            headers: {
              "Content-Type": "application/x-www-form-urlencoded",
            },
          }
        )
        .then((res) => {
          this.getTask();
          this.setState({
            task: "",
          });
          console.log(res);
        });
    }
  };

  getTask = () => {
    axios.get(endpoint + "/api/task").then((res) => {
      if (res.data) {
        this.setState({ items: res.data });
      } else {
        this.setState({ items: [] });
      }
    });
  };

  updateTask = (id) => {
    axios
      .put(
        endpoint + "/api/task/" + id,
        {},
        {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
        }
      )
      .then((res) => {
        this.getTask();
        console.log(res);
      });
  };

  undoTask = (id) => {
    axios
      .put(
        endpoint + "/api/undoTask/" + id,
        {},
        {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
        }
      )
      .then((res) => {
        this.getTask();
        console.log(res);
      });
  };

  deleteTask = (id) => {
    if (window.confirm("Are you sure you want to delete this task?")) {
      axios
        .delete(endpoint + "/api/deleteTask/" + id, {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
        })
        .then((res) => {
          console.log(res);
          this.getTask();
        })
        .catch((err) => console.log(err));
    }
  };

  deleteAllTasks = () => {
    if (window.confirm("Are you sure you want to delete all tasks?")) {
      axios
        .delete(endpoint + "/api/deleteAllTasks", {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
        })
        .then((res) => {
          console.log(res);
          this.getTask();
        })
        .catch((err) => console.log(err));
    }
  };
  

  renderItems = () => {
    return this.state.items.map((item) => {
      let color = "yellow";
      let style = {
        wordWrap: "break-word",
      };
      if (item.status) {
        color = "green";
        style["textDecorationLine"] = "line-through";
      }
      return (
        <Card key={item._id} color={color} fluid className="rough">
          <Card.Content style={{ backgroundColor: "#ebebeb" }}>
            <Card.Header textAlign="left">
              <div style={style}>{item.task}</div>
            </Card.Header>
            <Card.Meta textAlign="right">
              <Icon
                name="check circle"
                color="green"
                onClick={() => this.updateTask(item._id)}
              />
              <span style={{ paddingRight: 10 }}>Done</span>
              <Icon
                name="undo"
                color="orange"
                onClick={() => this.undoTask(item._id)}
              />
              <span style={{ paddingRight: 10 }}>Undo</span>
              <Icon
                name="delete"
                color="red"
                onClick={() => this.deleteTask(item._id)}
              />
              <span style={{ paddingRight: 10 }}>Delete</span>
            </Card.Meta>
          </Card.Content>
        </Card>
      );
    });
  };

  render() {
    const taskCount = this.state.items.length;
    return (
      <div>
        <div className="row">
          <Header className="header" as="h2" color="black">
            TO DO APP
          </Header>
        </div>
        <div className="row">
          <Form onSubmit={this.onSubmit} className="form-inline">
            <Input
              type="text"
              name="task"
              onChange={this.onChange}
              value={this.state.task}
              fluid
              placeholder="Add your new todo"
              className="input-task"
            />
            <Button type="submit" color="purple" icon style={{ marginLeft: "10px" }}>
              <Icon name="plus" />
            </Button>
          </Form>
        </div>
        <div className="row">
          <Card.Group>{this.renderItems()}</Card.Group>
        </div >
        <div className="row form-container">
          <p className="task-counter">You have {taskCount} pending tasks</p>
          <Button
            color="purple"
            style={{ marginLeft: "auto" }}
            onClick={this.deleteAllTasks}
          >
            Delete All
          </Button>
        </div>
        {/* <Button
            color="purple"
            style={{ marginLeft: "949px" }}
            onClick={this.deleteAllTasks}
          >
            Delete All
          </Button> */}
      </div>
    );
  }
}

export default ToDoList;