// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: coleta.proto

package coleta

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ContraCheque_Tipo int32

const (
	ContraCheque_MEMBRO   ContraCheque_Tipo = 0
	ContraCheque_SERVIDOR ContraCheque_Tipo = 1
)

// Enum value maps for ContraCheque_Tipo.
var (
	ContraCheque_Tipo_name = map[int32]string{
		0: "MEMBRO",
		1: "SERVIDOR",
	}
	ContraCheque_Tipo_value = map[string]int32{
		"MEMBRO":   0,
		"SERVIDOR": 1,
	}
)

func (x ContraCheque_Tipo) Enum() *ContraCheque_Tipo {
	p := new(ContraCheque_Tipo)
	*p = x
	return p
}

func (x ContraCheque_Tipo) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContraCheque_Tipo) Descriptor() protoreflect.EnumDescriptor {
	return file_coleta_proto_enumTypes[0].Descriptor()
}

func (ContraCheque_Tipo) Type() protoreflect.EnumType {
	return &file_coleta_proto_enumTypes[0]
}

func (x ContraCheque_Tipo) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContraCheque_Tipo.Descriptor instead.
func (ContraCheque_Tipo) EnumDescriptor() ([]byte, []int) {
	return file_coleta_proto_rawDescGZIP(), []int{4, 0}
}

type Remuneracao_Natureza int32

const (
	Remuneracao_R Remuneracao_Natureza = 0 //Receita
	Remuneracao_D Remuneracao_Natureza = 1 //Despesa
)

// Enum value maps for Remuneracao_Natureza.
var (
	Remuneracao_Natureza_name = map[int32]string{
		0: "R",
		1: "D",
	}
	Remuneracao_Natureza_value = map[string]int32{
		"R": 0,
		"D": 1,
	}
)

func (x Remuneracao_Natureza) Enum() *Remuneracao_Natureza {
	p := new(Remuneracao_Natureza)
	*p = x
	return p
}

func (x Remuneracao_Natureza) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Remuneracao_Natureza) Descriptor() protoreflect.EnumDescriptor {
	return file_coleta_proto_enumTypes[1].Descriptor()
}

func (Remuneracao_Natureza) Type() protoreflect.EnumType {
	return &file_coleta_proto_enumTypes[1]
}

func (x Remuneracao_Natureza) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Remuneracao_Natureza.Descriptor instead.
func (Remuneracao_Natureza) EnumDescriptor() ([]byte, []int) {
	return file_coleta_proto_rawDescGZIP(), []int{6, 0}
}

type Remuneracao_TipoReceita int32

const (
	Remuneracao_B Remuneracao_TipoReceita = 0 //Base
	Remuneracao_O Remuneracao_TipoReceita = 1 //Outros
)

// Enum value maps for Remuneracao_TipoReceita.
var (
	Remuneracao_TipoReceita_name = map[int32]string{
		0: "B",
		1: "O",
	}
	Remuneracao_TipoReceita_value = map[string]int32{
		"B": 0,
		"O": 1,
	}
)

func (x Remuneracao_TipoReceita) Enum() *Remuneracao_TipoReceita {
	p := new(Remuneracao_TipoReceita)
	*p = x
	return p
}

func (x Remuneracao_TipoReceita) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Remuneracao_TipoReceita) Descriptor() protoreflect.EnumDescriptor {
	return file_coleta_proto_enumTypes[2].Descriptor()
}

func (Remuneracao_TipoReceita) Type() protoreflect.EnumType {
	return &file_coleta_proto_enumTypes[2]
}

func (x Remuneracao_TipoReceita) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Remuneracao_TipoReceita.Descriptor instead.
func (Remuneracao_TipoReceita) EnumDescriptor() ([]byte, []int) {
	return file_coleta_proto_rawDescGZIP(), []int{6, 1}
}

// Estrutura com informações pertencentes a execução da coleta dos dados.
type ResultadoColeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Coleta   *Coleta           `protobuf:"bytes,1,opt,name=coleta,proto3" json:"coleta,omitempty"`
	Folha    *FolhaDePagamento `protobuf:"bytes,2,opt,name=folha,proto3" json:"folha,omitempty"`
	Procinfo *ProcInfo         `protobuf:"bytes,3,opt,name=procinfo,proto3" json:"procinfo,omitempty"`
}

func (x *ResultadoColeta) Reset() {
	*x = ResultadoColeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coleta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultadoColeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultadoColeta) ProtoMessage() {}

func (x *ResultadoColeta) ProtoReflect() protoreflect.Message {
	mi := &file_coleta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultadoColeta.ProtoReflect.Descriptor instead.
func (*ResultadoColeta) Descriptor() ([]byte, []int) {
	return file_coleta_proto_rawDescGZIP(), []int{0}
}

func (x *ResultadoColeta) GetColeta() *Coleta {
	if x != nil {
		return x.Coleta
	}
	return nil
}

func (x *ResultadoColeta) GetFolha() *FolhaDePagamento {
	if x != nil {
		return x.Folha
	}
	return nil
}

func (x *ResultadoColeta) GetProcinfo() *ProcInfo {
	if x != nil {
		return x.Procinfo
	}
	return nil
}

type ProcInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stdin  string   `protobuf:"bytes,1,opt,name=stdin,proto3" json:"stdin,omitempty"`    // String containing the standard input of the process.
	Stdout string   `protobuf:"bytes,2,opt,name=stdout,proto3" json:"stdout,omitempty"`  // String containing the standard output of the process.
	Stderr string   `protobuf:"bytes,3,opt,name=stderr,proto3" json:"stderr,omitempty"`  // String containing the standard error of the process.
	Cmd    string   `protobuf:"bytes,4,opt,name=cmd,proto3" json:"cmd,omitempty"`        // Command that has been executed
	CmdDir string   `protobuf:"bytes,5,opt,name=cmdDir,proto3" json:"cmdDir,omitempty"`  // Local directory, in which the command has been executed
	Status int32    `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"` // Exit code of the process executed
	Env    []string `protobuf:"bytes,7,rep,name=env,proto3" json:"env,omitempty"`        // Copy of strings representing the environment variables in the form ke=value
}

func (x *ProcInfo) Reset() {
	*x = ProcInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coleta_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcInfo) ProtoMessage() {}

func (x *ProcInfo) ProtoReflect() protoreflect.Message {
	mi := &file_coleta_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcInfo.ProtoReflect.Descriptor instead.
func (*ProcInfo) Descriptor() ([]byte, []int) {
	return file_coleta_proto_rawDescGZIP(), []int{1}
}

func (x *ProcInfo) GetStdin() string {
	if x != nil {
		return x.Stdin
	}
	return ""
}

func (x *ProcInfo) GetStdout() string {
	if x != nil {
		return x.Stdout
	}
	return ""
}

func (x *ProcInfo) GetStderr() string {
	if x != nil {
		return x.Stderr
	}
	return ""
}

func (x *ProcInfo) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

func (x *ProcInfo) GetCmdDir() string {
	if x != nil {
		return x.CmdDir
	}
	return ""
}

func (x *ProcInfo) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ProcInfo) GetEnv() []string {
	if x != nil {
		return x.Env
	}
	return nil
}

// Estrutura com informações pertencentes a coleta dos dados referentes a um órgão-mes-ano
type Coleta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChaveColeta        string                 `protobuf:"bytes,1,opt,name=chave_coleta,json=chaveColeta,proto3" json:"chave_coleta,omitempty"`
	Orgao              string                 `protobuf:"bytes,2,opt,name=orgao,proto3" json:"orgao,omitempty"`
	Mes                int32                  `protobuf:"varint,3,opt,name=mes,proto3" json:"mes,omitempty"`
	Ano                int32                  `protobuf:"varint,4,opt,name=ano,proto3" json:"ano,omitempty"`
	TimestampColeta    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp_coleta,json=timestampColeta,proto3" json:"timestamp_coleta,omitempty"`
	RepositorioColetor string                 `protobuf:"bytes,6,opt,name=repositorio_coletor,json=repositorioColetor,proto3" json:"repositorio_coletor,omitempty"`
	VersaoColetor      string                 `protobuf:"bytes,7,opt,name=versao_coletor,json=versaoColetor,proto3" json:"versao_coletor,omitempty"`
	DirColetor         string                 `protobuf:"bytes,8,opt,name=dir_coletor,json=dirColetor,proto3" json:"dir_coletor,omitempty"`
	Arquivos           []string               `protobuf:"bytes,9,rep,name=arquivos,proto3" json:"arquivos,omitempty"`
}

func (x *Coleta) Reset() {
	*x = Coleta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coleta_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coleta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coleta) ProtoMessage() {}

func (x *Coleta) ProtoReflect() protoreflect.Message {
	mi := &file_coleta_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coleta.ProtoReflect.Descriptor instead.
func (*Coleta) Descriptor() ([]byte, []int) {
	return file_coleta_proto_rawDescGZIP(), []int{2}
}

func (x *Coleta) GetChaveColeta() string {
	if x != nil {
		return x.ChaveColeta
	}
	return ""
}

func (x *Coleta) GetOrgao() string {
	if x != nil {
		return x.Orgao
	}
	return ""
}

func (x *Coleta) GetMes() int32 {
	if x != nil {
		return x.Mes
	}
	return 0
}

func (x *Coleta) GetAno() int32 {
	if x != nil {
		return x.Ano
	}
	return 0
}

func (x *Coleta) GetTimestampColeta() *timestamppb.Timestamp {
	if x != nil {
		return x.TimestampColeta
	}
	return nil
}

func (x *Coleta) GetRepositorioColetor() string {
	if x != nil {
		return x.RepositorioColetor
	}
	return ""
}

func (x *Coleta) GetVersaoColetor() string {
	if x != nil {
		return x.VersaoColetor
	}
	return ""
}

func (x *Coleta) GetDirColetor() string {
	if x != nil {
		return x.DirColetor
	}
	return ""
}

func (x *Coleta) GetArquivos() []string {
	if x != nil {
		return x.Arquivos
	}
	return nil
}

// Estrutura que faz referência a uma lista de contra-cheques
type FolhaDePagamento struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContraCheque []*ContraCheque `protobuf:"bytes,1,rep,name=contra_cheque,json=contraCheque,proto3" json:"contra_cheque,omitempty"`
}

func (x *FolhaDePagamento) Reset() {
	*x = FolhaDePagamento{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coleta_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FolhaDePagamento) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FolhaDePagamento) ProtoMessage() {}

func (x *FolhaDePagamento) ProtoReflect() protoreflect.Message {
	mi := &file_coleta_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FolhaDePagamento.ProtoReflect.Descriptor instead.
func (*FolhaDePagamento) Descriptor() ([]byte, []int) {
	return file_coleta_proto_rawDescGZIP(), []int{3}
}

func (x *FolhaDePagamento) GetContraCheque() []*ContraCheque {
	if x != nil {
		return x.ContraCheque
	}
	return nil
}

// Estrutura que faz referência a uma linha de contra-cheque em um órgão-mes-ano
type ContraCheque struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdContraCheque string            `protobuf:"bytes,1,opt,name=id_contra_cheque,json=idContraCheque,proto3" json:"id_contra_cheque,omitempty"`
	ChaveColeta    string            `protobuf:"bytes,2,opt,name=chave_coleta,json=chaveColeta,proto3" json:"chave_coleta,omitempty"`
	Nome           string            `protobuf:"bytes,3,opt,name=nome,proto3" json:"nome,omitempty"`
	Matricula      string            `protobuf:"bytes,4,opt,name=matricula,proto3" json:"matricula,omitempty"`
	Funcao         string            `protobuf:"bytes,5,opt,name=funcao,proto3" json:"funcao,omitempty"`
	LocalTrabalho  string            `protobuf:"bytes,6,opt,name=local_trabalho,json=localTrabalho,proto3" json:"local_trabalho,omitempty"`
	Tipo           ContraCheque_Tipo `protobuf:"varint,7,opt,name=tipo,proto3,enum=ContraCheque_Tipo" json:"tipo,omitempty"`
	Ativo          bool              `protobuf:"varint,8,opt,name=ativo,proto3" json:"ativo,omitempty"`
	Remuneracoes   *Remuneracoes     `protobuf:"bytes,9,opt,name=remuneracoes,proto3" json:"remuneracoes,omitempty"`
}

func (x *ContraCheque) Reset() {
	*x = ContraCheque{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coleta_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContraCheque) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContraCheque) ProtoMessage() {}

func (x *ContraCheque) ProtoReflect() protoreflect.Message {
	mi := &file_coleta_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContraCheque.ProtoReflect.Descriptor instead.
func (*ContraCheque) Descriptor() ([]byte, []int) {
	return file_coleta_proto_rawDescGZIP(), []int{4}
}

func (x *ContraCheque) GetIdContraCheque() string {
	if x != nil {
		return x.IdContraCheque
	}
	return ""
}

func (x *ContraCheque) GetChaveColeta() string {
	if x != nil {
		return x.ChaveColeta
	}
	return ""
}

func (x *ContraCheque) GetNome() string {
	if x != nil {
		return x.Nome
	}
	return ""
}

func (x *ContraCheque) GetMatricula() string {
	if x != nil {
		return x.Matricula
	}
	return ""
}

func (x *ContraCheque) GetFuncao() string {
	if x != nil {
		return x.Funcao
	}
	return ""
}

func (x *ContraCheque) GetLocalTrabalho() string {
	if x != nil {
		return x.LocalTrabalho
	}
	return ""
}

func (x *ContraCheque) GetTipo() ContraCheque_Tipo {
	if x != nil {
		return x.Tipo
	}
	return ContraCheque_MEMBRO
}

func (x *ContraCheque) GetAtivo() bool {
	if x != nil {
		return x.Ativo
	}
	return false
}

func (x *ContraCheque) GetRemuneracoes() *Remuneracoes {
	if x != nil {
		return x.Remuneracoes
	}
	return nil
}

// Estrutura que faz referência a uma lista de Remunerações
type Remuneracoes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Remuneracao []*Remuneracao `protobuf:"bytes,1,rep,name=remuneracao,proto3" json:"remuneracao,omitempty"`
}

func (x *Remuneracoes) Reset() {
	*x = Remuneracoes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coleta_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Remuneracoes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Remuneracoes) ProtoMessage() {}

func (x *Remuneracoes) ProtoReflect() protoreflect.Message {
	mi := &file_coleta_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Remuneracoes.ProtoReflect.Descriptor instead.
func (*Remuneracoes) Descriptor() ([]byte, []int) {
	return file_coleta_proto_rawDescGZIP(), []int{5}
}

func (x *Remuneracoes) GetRemuneracao() []*Remuneracao {
	if x != nil {
		return x.Remuneracao
	}
	return nil
}

// Estrutura que faz referência a uma linha de remuneração em um órgão-mes-ano
type Remuneracao struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Natureza    Remuneracao_Natureza    `protobuf:"varint,1,opt,name=natureza,proto3,enum=Remuneracao_Natureza" json:"natureza,omitempty"`
	Categoria   string                  `protobuf:"bytes,2,opt,name=categoria,proto3" json:"categoria,omitempty"`
	Item        string                  `protobuf:"bytes,3,opt,name=item,proto3" json:"item,omitempty"`
	Valor       float64                 `protobuf:"fixed64,4,opt,name=valor,proto3" json:"valor,omitempty"`
	TipoReceita Remuneracao_TipoReceita `protobuf:"varint,5,opt,name=tipo_receita,json=tipoReceita,proto3,enum=Remuneracao_TipoReceita" json:"tipo_receita,omitempty"`
}

func (x *Remuneracao) Reset() {
	*x = Remuneracao{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coleta_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Remuneracao) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Remuneracao) ProtoMessage() {}

func (x *Remuneracao) ProtoReflect() protoreflect.Message {
	mi := &file_coleta_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Remuneracao.ProtoReflect.Descriptor instead.
func (*Remuneracao) Descriptor() ([]byte, []int) {
	return file_coleta_proto_rawDescGZIP(), []int{6}
}

func (x *Remuneracao) GetNatureza() Remuneracao_Natureza {
	if x != nil {
		return x.Natureza
	}
	return Remuneracao_R
}

func (x *Remuneracao) GetCategoria() string {
	if x != nil {
		return x.Categoria
	}
	return ""
}

func (x *Remuneracao) GetItem() string {
	if x != nil {
		return x.Item
	}
	return ""
}

func (x *Remuneracao) GetValor() float64 {
	if x != nil {
		return x.Valor
	}
	return 0
}

func (x *Remuneracao) GetTipoReceita() Remuneracao_TipoReceita {
	if x != nil {
		return x.TipoReceita
	}
	return Remuneracao_B
}

var File_coleta_proto protoreflect.FileDescriptor

var file_coleta_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6c, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x82, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x64, 0x6f, 0x43, 0x6f, 0x6c,
	0x65, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x43, 0x6f, 0x6c, 0x65, 0x74, 0x61, 0x52, 0x06, 0x63, 0x6f,
	0x6c, 0x65, 0x74, 0x61, 0x12, 0x27, 0x0a, 0x05, 0x66, 0x6f, 0x6c, 0x68, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x46, 0x6f, 0x6c, 0x68, 0x61, 0x44, 0x65, 0x50, 0x61, 0x67,
	0x61, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x52, 0x05, 0x66, 0x6f, 0x6c, 0x68, 0x61, 0x12, 0x25, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x63, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x63,
	0x69, 0x6e, 0x66, 0x6f, 0x22, 0xa4, 0x01, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x63, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x64, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x74, 0x64, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6d, 0x64,
	0x44, 0x69, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6d, 0x64, 0x44, 0x69,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x76,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x22, 0xc1, 0x02, 0x0a, 0x06,
	0x43, 0x6f, 0x6c, 0x65, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x76, 0x65, 0x5f,
	0x63, 0x6f, 0x6c, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68,
	0x61, 0x76, 0x65, 0x43, 0x6f, 0x6c, 0x65, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x67,
	0x61, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x61, 0x6f, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6d, 0x65,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x6e, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x61, 0x6e, 0x6f, 0x12, 0x45, 0x0a, 0x10, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x5f, 0x63, 0x6f, 0x6c, 0x65, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x43, 0x6f, 0x6c, 0x65, 0x74, 0x61, 0x12, 0x2f, 0x0a, 0x13, 0x72, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6f, 0x5f, 0x63, 0x6f, 0x6c, 0x65, 0x74, 0x6f,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x6f, 0x43, 0x6f, 0x6c, 0x65, 0x74, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x76,
	0x65, 0x72, 0x73, 0x61, 0x6f, 0x5f, 0x63, 0x6f, 0x6c, 0x65, 0x74, 0x6f, 0x72, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6f, 0x43, 0x6f, 0x6c, 0x65, 0x74,
	0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x69, 0x72, 0x5f, 0x63, 0x6f, 0x6c, 0x65, 0x74, 0x6f,
	0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x69, 0x72, 0x43, 0x6f, 0x6c, 0x65,
	0x74, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x72, 0x71, 0x75, 0x69, 0x76, 0x6f, 0x73, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x61, 0x72, 0x71, 0x75, 0x69, 0x76, 0x6f, 0x73, 0x22,
	0x46, 0x0a, 0x10, 0x46, 0x6f, 0x6c, 0x68, 0x61, 0x44, 0x65, 0x50, 0x61, 0x67, 0x61, 0x6d, 0x65,
	0x6e, 0x74, 0x6f, 0x12, 0x32, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x5f, 0x63, 0x68,
	0x65, 0x71, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x43, 0x68, 0x65, 0x71, 0x75, 0x65, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x43, 0x68, 0x65, 0x71, 0x75, 0x65, 0x22, 0xdf, 0x02, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x43, 0x68, 0x65, 0x71, 0x75, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x64, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x5f, 0x63, 0x68, 0x65, 0x71, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x69, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x43, 0x68, 0x65, 0x71,
	0x75, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6c, 0x65,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x76, 0x65, 0x43,
	0x6f, 0x6c, 0x65, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x74,
	0x72, 0x69, 0x63, 0x75, 0x6c, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61,
	0x74, 0x72, 0x69, 0x63, 0x75, 0x6c, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x75, 0x6e, 0x63, 0x61,
	0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x75, 0x6e, 0x63, 0x61, 0x6f, 0x12,
	0x25, 0x0a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x74, 0x72, 0x61, 0x62, 0x61, 0x6c, 0x68,
	0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x54, 0x72,
	0x61, 0x62, 0x61, 0x6c, 0x68, 0x6f, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x43, 0x68, 0x65,
	0x71, 0x75, 0x65, 0x2e, 0x54, 0x69, 0x70, 0x6f, 0x52, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x12, 0x14,
	0x0a, 0x05, 0x61, 0x74, 0x69, 0x76, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61,
	0x74, 0x69, 0x76, 0x6f, 0x12, 0x31, 0x0a, 0x0c, 0x72, 0x65, 0x6d, 0x75, 0x6e, 0x65, 0x72, 0x61,
	0x63, 0x6f, 0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x52, 0x65, 0x6d,
	0x75, 0x6e, 0x65, 0x72, 0x61, 0x63, 0x6f, 0x65, 0x73, 0x52, 0x0c, 0x72, 0x65, 0x6d, 0x75, 0x6e,
	0x65, 0x72, 0x61, 0x63, 0x6f, 0x65, 0x73, 0x22, 0x20, 0x0a, 0x04, 0x54, 0x69, 0x70, 0x6f, 0x12,
	0x0a, 0x0a, 0x06, 0x4d, 0x45, 0x4d, 0x42, 0x52, 0x4f, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x53,
	0x45, 0x52, 0x56, 0x49, 0x44, 0x4f, 0x52, 0x10, 0x01, 0x22, 0x3e, 0x0a, 0x0c, 0x52, 0x65, 0x6d,
	0x75, 0x6e, 0x65, 0x72, 0x61, 0x63, 0x6f, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x0b, 0x72, 0x65, 0x6d,
	0x75, 0x6e, 0x65, 0x72, 0x61, 0x63, 0x61, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x52, 0x65, 0x6d, 0x75, 0x6e, 0x65, 0x72, 0x61, 0x63, 0x61, 0x6f, 0x52, 0x0b, 0x72, 0x65,
	0x6d, 0x75, 0x6e, 0x65, 0x72, 0x61, 0x63, 0x61, 0x6f, 0x22, 0xfc, 0x01, 0x0a, 0x0b, 0x52, 0x65,
	0x6d, 0x75, 0x6e, 0x65, 0x72, 0x61, 0x63, 0x61, 0x6f, 0x12, 0x31, 0x0a, 0x08, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x7a, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x52, 0x65,
	0x6d, 0x75, 0x6e, 0x65, 0x72, 0x61, 0x63, 0x61, 0x6f, 0x2e, 0x4e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x7a, 0x61, 0x52, 0x08, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x7a, 0x61, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x74,
	0x65, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x6f, 0x72, 0x12, 0x3b, 0x0a, 0x0c, 0x74, 0x69, 0x70, 0x6f, 0x5f, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x52, 0x65, 0x6d,
	0x75, 0x6e, 0x65, 0x72, 0x61, 0x63, 0x61, 0x6f, 0x2e, 0x54, 0x69, 0x70, 0x6f, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x74, 0x61, 0x52, 0x0b, 0x74, 0x69, 0x70, 0x6f, 0x52, 0x65, 0x63, 0x65, 0x69, 0x74,
	0x61, 0x22, 0x18, 0x0a, 0x08, 0x4e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x7a, 0x61, 0x12, 0x05, 0x0a,
	0x01, 0x52, 0x10, 0x00, 0x12, 0x05, 0x0a, 0x01, 0x44, 0x10, 0x01, 0x22, 0x1b, 0x0a, 0x0b, 0x54,
	0x69, 0x70, 0x6f, 0x52, 0x65, 0x63, 0x65, 0x69, 0x74, 0x61, 0x12, 0x05, 0x0a, 0x01, 0x42, 0x10,
	0x00, 0x12, 0x05, 0x0a, 0x01, 0x4f, 0x10, 0x01, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x64, 0x6f, 0x73, 0x6a, 0x75, 0x73, 0x62,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6c, 0x65, 0x74, 0x61, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coleta_proto_rawDescOnce sync.Once
	file_coleta_proto_rawDescData = file_coleta_proto_rawDesc
)

func file_coleta_proto_rawDescGZIP() []byte {
	file_coleta_proto_rawDescOnce.Do(func() {
		file_coleta_proto_rawDescData = protoimpl.X.CompressGZIP(file_coleta_proto_rawDescData)
	})
	return file_coleta_proto_rawDescData
}

var file_coleta_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_coleta_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_coleta_proto_goTypes = []interface{}{
	(ContraCheque_Tipo)(0),        // 0: ContraCheque.Tipo
	(Remuneracao_Natureza)(0),     // 1: Remuneracao.Natureza
	(Remuneracao_TipoReceita)(0),  // 2: Remuneracao.TipoReceita
	(*ResultadoColeta)(nil),       // 3: ResultadoColeta
	(*ProcInfo)(nil),              // 4: ProcInfo
	(*Coleta)(nil),                // 5: Coleta
	(*FolhaDePagamento)(nil),      // 6: FolhaDePagamento
	(*ContraCheque)(nil),          // 7: ContraCheque
	(*Remuneracoes)(nil),          // 8: Remuneracoes
	(*Remuneracao)(nil),           // 9: Remuneracao
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_coleta_proto_depIdxs = []int32{
	5,  // 0: ResultadoColeta.coleta:type_name -> Coleta
	6,  // 1: ResultadoColeta.folha:type_name -> FolhaDePagamento
	4,  // 2: ResultadoColeta.procinfo:type_name -> ProcInfo
	10, // 3: Coleta.timestamp_coleta:type_name -> google.protobuf.Timestamp
	7,  // 4: FolhaDePagamento.contra_cheque:type_name -> ContraCheque
	0,  // 5: ContraCheque.tipo:type_name -> ContraCheque.Tipo
	8,  // 6: ContraCheque.remuneracoes:type_name -> Remuneracoes
	9,  // 7: Remuneracoes.remuneracao:type_name -> Remuneracao
	1,  // 8: Remuneracao.natureza:type_name -> Remuneracao.Natureza
	2,  // 9: Remuneracao.tipo_receita:type_name -> Remuneracao.TipoReceita
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_coleta_proto_init() }
func file_coleta_proto_init() {
	if File_coleta_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coleta_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultadoColeta); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coleta_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coleta_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coleta); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coleta_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FolhaDePagamento); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coleta_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContraCheque); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coleta_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Remuneracoes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coleta_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Remuneracao); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_coleta_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_coleta_proto_goTypes,
		DependencyIndexes: file_coleta_proto_depIdxs,
		EnumInfos:         file_coleta_proto_enumTypes,
		MessageInfos:      file_coleta_proto_msgTypes,
	}.Build()
	File_coleta_proto = out.File
	file_coleta_proto_rawDesc = nil
	file_coleta_proto_goTypes = nil
	file_coleta_proto_depIdxs = nil
}
