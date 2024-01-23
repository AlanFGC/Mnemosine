import { useCallback, useEffect } from 'react';
import { useLexicalComposerContext } from '@lexical/react/LexicalComposerContext';
import Debounce from '../../../utils/Debounce';

type LocalStoragePluginProps = {
  namespace: string;
};

const DEBOUNCETIME = 300;

function LocalStoragePlugin({ namespace }: LocalStoragePluginProps): React.ReactNode {
  const [editor] = useLexicalComposerContext();

  // SAVE CONTENT TO LOCAL STORAGE
  const saveContent = useCallback(
    (content: string) => {
      localStorage.setItem(namespace, content);
    },
    [namespace],
  );

  // Debounce to save content to local storage
  const debounceFunc = Debounce(saveContent, DEBOUNCETIME);

  // trigger debounce whenever the tree or the content changes.
  useEffect(() => (
    editor.registerUpdateListener(
      ({ editorState, dirtyElements, dirtyLeaves }) => {
        if (dirtyElements.size === 0 && dirtyLeaves.size === 0) return;

        const serializeState = JSON.stringify(editorState);
        debounceFunc(serializeState);
      },
    )
  ), [debounceFunc, editor]);

  return null;
}
export default LocalStoragePlugin;
