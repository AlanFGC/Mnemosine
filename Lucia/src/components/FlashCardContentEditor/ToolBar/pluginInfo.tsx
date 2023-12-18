import { QuestionMarkCircleIcon } from '@heroicons/react/20/solid';

export type EventType = string;

export const eventTypes: { [key: string]: EventType } = {
  paragraph: 'paragraph',
  h1: 'h1',
  h2: 'h2',
  ul: 'ul',
  ol: 'ol',
  quote: 'quote',
  formatCode: 'formatCode',
  formatUndo: 'formatUndo',
  formatRedo: 'formatRedo',
  formatBold: 'formatBold',
  formatItalic: 'formatItalic',
  formatUnderline: 'formatUnderline',
  formatStrike: 'formatStrike',
  formatInsertLink: 'formatInsertLink',
  formatAlignLeft: 'formatAlignLeft',
  formatAlignCenter: 'formatAlignCenter',
  formatAlignRight: 'formatAlignRight',
  insertImage: 'insertImage',
  insertToken: 'insertToken',
};

export interface FlashContEditorPlugin {
  id: number,
  Icon: React.ForwardRefExoticComponent<
  React.PropsWithoutRef<React.SVGProps<SVGSVGElement>> &
  { title?: string, titleId?: string } &
  React.RefAttributes<SVGSVGElement>
  >,
  event: EventType,
}

export const pluginsList: FlashContEditorPlugin[] = [
  // Create AnswerToken
  {
    id: 1,
    Icon: QuestionMarkCircleIcon,
    event: eventTypes.insertToken,
  },
];

export default pluginsList;
