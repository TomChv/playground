import React, { useState } from "react";

import List from "./components/List";
import TaskForm from "./components/TaskForm";

import "./index.css";

import { NoteRo } from "./dto/apiDto";

function App(): JSX.Element {
  const [taskList, setTaskList] = useState<NoteRo[]>([]);
  const [isAdding, setIsAdding] = useState<boolean>(false);

  return (
    <div className="app">
      <div className="title">
        <div className="empty"></div>
        <p>Nest Todolist powered by React</p>
        <button className="buttonAdd" onClick={() => setIsAdding(!isAdding)}>
          +
        </button>
      </div>
      {isAdding ? (
        <TaskForm
          taskList={taskList}
          setTaskList={setTaskList}
          setIsAdding={setIsAdding}
        />
      ) : (
        <div className="taskPanel">
          <List taskList={taskList} setTaskList={setTaskList} />
        </div>
      )}
    </div>
  );
}

export default App;
