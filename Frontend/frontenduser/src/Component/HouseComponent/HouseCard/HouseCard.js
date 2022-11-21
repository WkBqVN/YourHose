import React from 'react';
import { Card, Button } from 'react-bootstrap';
import PropTypes from 'prop-types';
import "./HouseCard.css";

function HouseCard() {
    let cardList = getCardList()
    return (
        [
            <div>
                <GenerateHouseCard />
            </div>
        ]
    )
}
function GenerateHouseCard({ img,price,comment,address }) {
    <div className="HouseCard">
        <Card>
            <Card.Img variant="top" src={img} />
            <Card.Body>
                <Card.Title>{price}</Card.Title>
                <Card.Text>
                    {comment}
                </Card.Text>
                <Card.Title>{address}</Card.Title>
                <Button variant="primary">Read More</Button>
            </Card.Body>
        </Card>
    </div>
}
function getCardList() {

}

HouseCard.propTypes = {
    img : PropTypes.string.isRequired,
    address : PropTypes.string.isRequired,
    comment : PropTypes.string.isRequired,
    price : PropTypes.number.isRequired,
}
export default HouseCard;  
