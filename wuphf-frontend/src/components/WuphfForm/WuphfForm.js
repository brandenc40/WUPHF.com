import axios from 'axios';
import React from 'react';
import ClipLoader from 'react-spinners/ClipLoader';
import wuphfSentAudioPath from '../../resources/audio/wuphf_wuphfalert.wav';
import wuphfLogo from '../../resources/images/logo.png';
import './WuphfForm.css';

const wuphfSentAudio = new Audio(wuphfSentAudioPath);

export default class WuphfForm extends React.Component {
  constructor(props) {
    super(props);

    var formFields = {};
    this.props.formOptions.forEach(option => {
      formFields[option.key] = '';
    });

    this.state = {
      formFields: formFields,
      errorMsg: '',
      loading: false
    };
  }

  onFormSubmit(e) {
    e.preventDefault();

    this.setState({
      errorMsg: '',
      loading: true
    });

    var formData = new FormData();
    this.props.formOptions.forEach(option => {
      formData.append(option.key, this.state.formFields[option.key]);
    });

    axios
      .post('/api/wuphf', formData)
      .then(_ => {
        wuphfSentAudio.play();
        this.setState({
          loading: false
        });
      })
      .catch(error => {
        this.setState({
          errorMsg: error.response.data,
          loading: false
        });
      });
  }

  handleChange(e) {
    let formFields = { ...this.state.formFields };
    formFields[e.target.name] = e.target.value;

    this.setState({
      formFields
    });
  }

  renderFormOptions() {
    const { formOptions } = this.props;

    let formUI = formOptions.map((option, i) => (
      <div key={i} className='wuphf-form-input-div'>
        {option.inputType === 'textarea' ? (
          <textarea
            className='wuphf-form-textarea'
            id={option.key}
            name={option.key}
            placeholder={option.placeholder}
            onChange={this.handleChange.bind(this)}
            value={this.state.formFields[option.key]}
            required={option.required}
          />
        ) : (
          <input
            className='wuphf-form-input'
            type={option.inputType}
            id={option.key}
            name={option.key}
            placeholder={option.placeholder}
            onChange={this.handleChange.bind(this)}
            value={this.state.formFields[option.key]}
            required={option.required}
          />
        )}
      </div>
    ));
    return formUI;
  }

  render() {
    return (
      <div className='wuphf-form'>
        <img alt='WUPHF Logo' src={wuphfLogo} className='logo-image' />
        <form onSubmit={this.onFormSubmit.bind(this)}>
          {this.renderFormOptions()}
          <button className='wuphf-submit-button'>Send Wuphf!</button>
        </form>
        <div className='wuphf-error'>{this.state.errorMsg}</div>
        <div className='loading-spinner'>
          <ClipLoader
            size={100}
            color={'#524A93'}
            loading={this.state.loading}
          />
        </div>
      </div>
    );
  }
}
