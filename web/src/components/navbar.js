import React, { Fragment } from 'react';
// import styled from '@emotion/core';
import styled from '@emotion/styled'


export default function PageContainer(props) {
    return (
        <Fragment>
            <nav>
                <NavWrapper>
                    <LeftSide>
                        <BrandWrapper>
                            <div>WasteChain</div>
                        </BrandWrapper>
                    </LeftSide>
                    <RightSide>
                        <LinkWrapper>
                            <p>Link</p>
                        </LinkWrapper>
                        <LinkWrapper>
                            <p>Link</p>
                        </LinkWrapper>
                    </RightSide>
                </NavWrapper>
            </nav>
        </Fragment>
    )
}

const NavWrapper = styled.div`
    display: flex;
    justify-content: space-between;
    padding: 38px;
    background-color: red;
`

const LeftSide = styled.div`

`

const BrandWrapper = styled.div`
    font-size: 32px;
`

const RightSide = styled.div`
    display: flex;
`

const LinkWrapper = styled.div`
    height: 22px;
    border-bottom: 1px solid transparent;
    transition: border-bottom 0.5s;
`
const NavLink = styled.a`
    color: #8a8a8a;
    text-decoration: none;
    transition: color 0.5s;
`
