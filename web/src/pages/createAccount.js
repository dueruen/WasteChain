import { Mutation } from 'react-apollo'
import gql from 'graphql-tag';
import React, { Component } from 'react'

/**
 * Mutation to create an employee
 */
const CREATE_EMPLOYEE =
gql`
    mutation CreateEmployee($authData: AuthData!,$name: String!, $companyID: String!)
    {createEmployee(employee:
        {
            authData: $authData,
            name: $name,
            companyID: $companyID,
        }
    ){id}}
`;

class CreateEmployeePage extends Component {
    state = {
      userName: '',
      password: '',
      companyID: '',
      name: ''
    }

    render() {
      const { userName, password, companyID, name } = this.state
      return (
        <section>
          <h2>Create Employee</h2>
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
                Company ID
                <input
                value={companyID}
                onChange={e => this.setState({ companyID: e.target.value })}
                type="text"
                required
                />
            </label>
            <label>
                Name of Employee
                <input
                value={name}
                onChange={e => this.setState({ name: e.target.value })}
                type="text"
                required
                />
            </label>
          </form>
          <Mutation mutation={CREATE_EMPLOYEE} variables={{ authData : {userName, password}, companyID, name }}>
              {createEmployee => <button onClick={createEmployee}>Create Employee</button>}
          </Mutation>
        </section>
      )
    }
  }

  export default CreateEmployeePage
