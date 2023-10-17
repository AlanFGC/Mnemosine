import { Button } from 'antd';

function sayHello(): void {
  console.log('Hello World!');
}

function App() {
  return (
    <div className="App">
      <h1>Welcome to Mnesomine</h1>
      <Button type="primary" onClick={sayHello}>Button</Button>
    </div>
  );
}

export default App;
