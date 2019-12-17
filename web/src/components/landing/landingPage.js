import React, { Fragment } from 'react';
import styled from '@emotion/styled'
import AllShipmentsList from '../shipments/allShipments'


export default function PageContainer(props) {
    return (
        <section>
            <Fragment>
                <h1>LANDING PAGE</h1>
                <AllShipmentsList></AllShipmentsList>
            </Fragment>
        </section>
    )
}
