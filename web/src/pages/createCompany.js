import { Mutation } from 'react-apollo'
import gql from 'graphql-tag';
import React, { Component } from 'react'

/**
 * Mutation to create an company
 */
const CREATE_COMPANY =
gql`
    mutation CreateCompany($authData: AuthData!,$name: String!, $address: CreateAddress!, $contactinfo: CreateContactInfo!)
    {createCompany(company:
        {
            authData: $authData,
            name: $name,
            address: $address,
            contactinfo: $contactinfo
        }
    ){id}}
`;

class CreateCompanyPage extends Component {
    state = {
      userName: '',
      password: '',
      name: '',
      roadName: '',
      number: 0,
      ZipCode: 0,
      title: 'Main contact',
      phoneNumber: 0,
      mail: '',
    }

    render() {
      const { userName, password, name, roadName, number, ZipCode, title, phoneNumber, mail} = this.state;
      return (
        <section>
          <h2>Create Company</h2>
          <form>
            <label>
                Username
                <input
                value={userName}
                onChange={e => this.setState({ userName: e.target.value })}
                type="text"
                required
                />
            </label>
            <label>
                Password
                <input
                value={password}
                onChange={e => this.setState({ password: e.target.value })}
                type="password"
                required
                />
            </label>
            <label>
                Name of Company
                <input
                value={name}
                onChange={e => this.setState({ name: e.target.value })}
                type="text"
                required
                />
            </label>
            <label>
                Road name
                <input
                value={roadName}
                onChange={e => this.setState({ roadName: e.target.value })}
                type="text"
                required
                />
            </label>
            <label>
                Road number
                <input
                value={number}
                onChange={e => this.setState({ number: e.target.value })}
                type="text"
                required
                />
            </label>
            <label>
                Zip Code
                <input
                value={ZipCode}
                onChange={e => this.setState({ ZipCode: e.target.value })}
                type="text"
                required
                />
            </label>
            <label>
                PhoneNumber
                <input
                value={phoneNumber}
                onChange={e => this.setState({ phoneNumber: e.target.value })}
                type="number"
                required
                />
            </label>
            <label>
                Mail
                <input
                value={mail}
                onChange={e => this.setState({ mail: e.target.value })}
                type="text"
                required
                />
            </label>
          </form>
          <Mutation mutation={CREATE_COMPANY} variables={{ authData : {userName, password}, name, address: {roadName, number, ZipCode}, contactinfo: {title, phoneNumber, mail} }}>
              {createCompany => <button onClick={createCompany}>Create Company</button>}
          </Mutation>
        </section>
      )
    }
  }

  export default CreateCompanyPage
