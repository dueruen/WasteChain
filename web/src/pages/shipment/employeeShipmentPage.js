import React, {Component} from 'react'
import ShipmentLink from '../../components/shipments/shipmentLink/shipmentLink'
import gql from "graphql-tag";
import { Query } from 'react-apollo'
import ReactLoading from 'react-loading';


/**
 * Query to list all shipments
 */
const LIST_USERS_SHIPMENTS =
    gql`
    query ListUsersShipments($userID: String!) {listUsersShipments(userID: $userID) {
    id
    }}
`;

class EmployeeShipmentPage extends Component {
    state = {
        userID: JSON.parse(localStorage.getItem('me'))["id"]
      }
    render() {
        const { userID } = this.state
        return(
            <Query query={LIST_USERS_SHIPMENTS} variables={{userID}}>
                {({ loading, error, data }) => {
                    if (loading) return <ReactLoading type={'spin'}color={'#8bb849'} height={'20%'} width={'20%'}/>
                    if (error) return <div>You're probably not logged into a user account</div>
                    const shipmentsToRender = data.listUsersShipments

                    return (
                        <section>
                            <h2>My Shipments</h2>
                            {shipmentsToRender.slice(0).reverse().map(shipment => <ShipmentLink key={shipment.id} id={shipment.id}/>)}
                        </section>
                    )
                }}
            </Query>
        )
    }
}

export default EmployeeShipmentPage
