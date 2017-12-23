// @flow

import AppBar from 'material-ui/AppBar'
import { Card, CardHeader, CardText } from 'material-ui/Card'
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider'
import React, { Component } from 'react'

import './App.css'
import gravatarUrl from './gravatar-url'

type Props = {}

type Message = {
  name: string,
  email: string,
  date: string,
  text: string
}

type State = {
  messages: Message[]
}

const formatDate = epoch => {
  const date = new Date(0)
  date.setUTCSeconds(epoch)
  const monthNames = [
    'January',
    'February',
    'March',
    'April',
    'May',
    'June',
    'July',
    'August',
    'September',
    'October',
    'November',
    'December'
  ]

  var day = date.getDate()
  var monthIndex = date.getMonth()
  var year = date.getFullYear()

  return day + ' ' + monthNames[monthIndex] + ' ' + year
}

class App extends Component<Props, State> {
  constructor(props: Props) {
    super(props)
    this.state = {
      messages: []
    }
  }
  componentDidMount() {
    fetch('/api/v1/messages')
      .then(res => {
        if (res.status !== 200) {
          console.log('response err', res)
        }
        return res.json()
      })
      .then(json => {
        this.setState({ messages: json.data })
      })
      .catch(err => {
        console.log('fetch err', err)
      })
  }
  render() {
    return (
      <MuiThemeProvider>
        <div>
          <AppBar title="Robo Retro" showMenuIconButton={false} />
          {this.state.messages.map((msg, i) => (
            <Card expanded key={i}>
              <CardHeader
                title={msg.name}
                subtitle={formatDate(msg.date)}
                avatar={gravatarUrl(msg.email)}
              />
              <CardText>{msg.text}</CardText>
            </Card>
          ))}
        </div>
      </MuiThemeProvider>
    )
  }
}

export default App
