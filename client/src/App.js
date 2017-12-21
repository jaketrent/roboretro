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

class App extends Component<Props, State> {
  constructor(props: Props) {
    super(props)
    this.state = {
      messages: [
        {
          name: 'Jakeroo',
          email: 'trent.jake@gmail.com',
          date: '27 Dec 2017',
          text: `
            Key players will take ownership of their innovations by virtually
            growing mobile stakeholders. We aim to globally virtualise our
            architecture by ethically reusing our immersive end-to-end user
            experiences.
`
        }
      ]
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
                subtitle={msg.date}
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
