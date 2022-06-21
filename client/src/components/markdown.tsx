import MarkdownIt from 'markdown-it'
import markdownItPrism from 'markdown-it-prism'

type Props = {
  content: string
}
export const MarkDown = (props: Props) => {
  const testText = '*test* text'

  const md = new MarkdownIt()
  md.use(markdownItPrism)

  const renderedMd = md.render(testText)

  return <div dangerouslySetInnerHTML={{ __html: renderedMd }} />
}
