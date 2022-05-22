import { Button } from '@mui/material'
import type { ButtonProps } from '@mui/material'
import Link, { LinkProps } from 'next/link'

type Props = {
  href: LinkProps['href']
  text: string
  buttonProps?: ButtonProps
}

const LinkedButton = (props: Props) => {
  return (
    <Link href={props.href} passHref>
      <Button {...props.buttonProps}>{props.text}</Button>
    </Link>
  )
}

export default LinkedButton
