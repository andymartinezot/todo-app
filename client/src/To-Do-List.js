import React, {Component} from "react";
import axios from "axios"; //It makes the API calls to the backend
import {Card, Header, Form, Input, Icon} from "semantic-ui-react";

let endpoint = "http://localhost:9000";

class ToDoList extends Component{
    constructor(props){
        super(props);

        this.state ={
            task:"",
            items:[],
        };
    }
    ComponentDidMount(){
        this.getTask();
    }
    render(){
        return(
            <div>
                <div className="row">
                    <Header className = "header" as="h2" color="yellow">
                        TO DO LIST
                    </Header>
                </div>
                <div className="row">
                    <Form onSubmit={this.onSubmit}>
                        <Input
                        type="text"
                        name="task"
                        onChange={this.onChange}
                        value={this.state.task}
                        fluid
                        placeholder="Create Task"
                        />
                    </Form>
                </div>
            </div>
        );
    }
}

export default ToDoList;