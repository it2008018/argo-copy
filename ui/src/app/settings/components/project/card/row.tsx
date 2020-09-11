import * as React from 'react';
import {GetProp, SetProp} from '../../utils';
import {Banner, BannerIcon, BannerType} from '../banner/banner';
import {ArgoField, FieldData, FieldTypes, FieldValue} from './field';

interface CardRowProps<T> {
    fields: FieldData[];
    data: T | FieldValue;
    remove: () => void;
    save: (value: T | FieldValue) => Promise<any>;
    selected: boolean;
    toggleSelect: () => void;
    changed: boolean;
    onChange: (value: T | FieldValue) => void;
}

export class CardRow<T> extends React.Component<CardRowProps<T>> {
    get disabled(): boolean {
        if (!this.props.data) {
            return true;
        }
        if (Object.keys(this.props.data).length < this.props.fields.length) {
            return true;
        }
        for (const key of Object.keys(this.props.data)) {
            const cur = GetProp(this.props.data as T, key as keyof T).toString();
            if (cur === '' || cur === null) {
                return true;
            }
        }
        return false;
    }
    get dataIsFieldValue(): boolean {
        return this.isFieldValue(this.props.data);
    }
    get fieldsSetToAll(): string[] {
        if (this.dataIsFieldValue) {
            const field = this.props.fields[0];
            const comp = field.type === FieldTypes.ResourceKindSelector ? 'ANY' : '*';
            return this.props.data.toString() === comp ? [field.name] : [];
        }
        const fields = [];
        for (const key of Object.keys(this.props.data)) {
            if (GetProp(this.props.data as T, key as keyof T).toString() === '*') {
                fields.push(key);
            }
        }
        return fields;
    }
    constructor(props: CardRowProps<T>) {
        super(props);
        this.state = {
            data: this.props.data
        };
    }

    public render() {
        let update = this.dataIsFieldValue
            ? (value: FieldValue, _: keyof T) => {
                  this.props.onChange(value);
              }
            : (value: FieldValue, field: keyof T) => {
                  const change = {...(this.props.data as T)};
                  SetProp(change, field, value);
                  this.props.onChange(change);
              };
        update = update.bind(this);
        return (
            <div>
                <div className='card__input-container card__row'>
                    <div className='card__col-round-button card__col'>
                        <button
                            className={`project__button project__button-round project__button-select${this.props.selected ? '--selected' : ''}`}
                            onClick={this.props.toggleSelect}>
                            <i className='fa fa-check' />
                        </button>
                    </div>
                    {this.props.fields.map((field, i) => {
                        const curVal = this.dataIsFieldValue ? this.props.data.toString() : GetProp(this.props.data as T, field.name as keyof T).toString();
                        return (
                            <div key={field.name} className={`card__col-input card__col card__col-${field.size}`}>
                                <ArgoField field={field} onChange={val => update(val, field.name as keyof T)} data={curVal} />
                            </div>
                        );
                    })}
                    <div className='card__col-button card__col'>
                        <button
                            className={`project__button project__button-${this.props.selected ? 'error' : this.disabled ? 'disabled' : this.props.changed ? 'save' : 'saved'}`}
                            onClick={() => (this.props.selected ? this.props.remove() : this.disabled ? null : this.props.save(this.props.data))}>
                            {this.props.selected ? 'DELETE' : this.disabled ? 'EMPTY' : this.props.changed ? 'SAVE' : 'SAVED'}
                        </button>
                    </div>
                </div>
                {this.fieldsSetToAll.length > 0 ? this.allNoticeBanner(this.fieldsSetToAll) : null}
            </div>
        );
    }
    private allNoticeBanner(fields: string[]) {
        let fieldList: string = fields[0] + 's';
        fields.splice(0, 1);
        if (fields.length > 0) {
            const last = fields.pop();
            if (fields.length > 0) {
                for (const field of fields) {
                    fieldList += ', ' + field + 's';
                }
            }
            fieldList += ' and ' + last + 's';
        }

        return (
            <div className='card__row'>
                <div className='card__col-round-button card__col' />
                <div className={'card__col card__col-grow'}>{Banner(BannerType.Info, BannerIcon.Info, `Note: ${fieldList} are set to wildcard (*)`)}</div>
                <div className='card__col-button card__col' />
            </div>
        );
    }
    private isFieldValue(value: T | FieldValue): value is FieldValue {
        if ((typeof value as FieldValue) === 'string') {
            return true;
        }
        return false;
    }
}
