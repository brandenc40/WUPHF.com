import React from 'react';

export default class PaypalDonate extends React.Component {
  render() {
    return (
      <div
        style={{
          padding: '10px'
        }}
      >
        <p
          style={{
            color: 'white',
            maxWidth: '9vw',
            fontSize: '0.8rem'
          }}
        >
          Each WUPHF costs me $, any donations welcomed :)
        </p>
        <form
          action='https://www.paypal.com/cgi-bin/webscr'
          method='post'
          target='_top'
        >
          <input type='hidden' name='cmd' value='_donations' />
          <input type='hidden' name='business' value='2ST2DPURPKP5U' />
          <input type='hidden' name='currency_code' value='USD' />
          <input
            type='image'
            src='https://www.paypalobjects.com/en_US/i/btn/btn_donate_SM.gif'
            border='0'
            name='submit'
            title='PayPal - The safer, easier way to pay online!'
            alt='Donate with PayPal button'
          />
          <img
            alt=''
            border='0'
            src='https://www.paypal.com/en_US/i/scr/pixel.gif'
            width='1'
            height='1'
          />
        </form>
      </div>
    );
  }
}
