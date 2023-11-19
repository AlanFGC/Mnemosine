export default function SingleAnswer(): JSX.Element {
  return (
    <div>
      <h1>Unique answer:</h1>
      <br />
      <input type="text" name="answer" />
      <span>Explanation:</span>
      <input type="text" name="explanation" />
    </div>
  );
}
