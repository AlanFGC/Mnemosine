import { pluginsList } from './pluginInfo';

function Toolbar(): JSX.Element {
  return (
    <ul>
      {pluginsList.map((plugin) => (
        <li key={plugin.id}>
          <plugin.Icon />
          <span>Add Question</span>
        </li>
      ))}
    </ul>
  );
}

export default Toolbar;
