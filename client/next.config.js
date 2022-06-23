const { withGlobalCss } = require('next-global-css')
/** @type {import('next').NextConfig} */

const withConfig = withGlobalCss()
const nextConfig = withConfig({
  reactStrictMode: true,
})

module.exports = nextConfig
