/**
 * @file    ast_exp.h
 * @copyright defined in aergo/LICENSE.txt
 */

#ifndef _AST_EXP_H
#define _AST_EXP_H

#include "common.h"

#include "ast.h"
#include "ast_meta.h"
#include "ast_val.h"

#define exp_is_null(exp)            ((exp)->kind == EXP_NULL)
#define exp_is_lit(exp)             ((exp)->kind == EXP_LIT)
#define exp_is_type(exp)            ((exp)->kind == EXP_TYPE)
#define exp_is_id(exp)              ((exp)->kind == EXP_ID)
#define exp_is_array(exp)           ((exp)->kind == EXP_ARRAY)
#define exp_is_op(exp)              ((exp)->kind == EXP_OP)
#define exp_is_access(exp)          ((exp)->kind == EXP_ACCESS)
#define exp_is_call(exp)            ((exp)->kind == EXP_CALL)
#define exp_is_sql(exp)             ((exp)->kind == EXP_SQL)
#define exp_is_ternary(exp)         ((exp)->kind == EXP_TERNARY)
#define exp_is_tuple(exp)           ((exp)->kind == EXP_TUPLE)

#define exp_can_be_lval(exp)                                                   \
    (!meta_is_const(&(exp)->meta) &&                                           \
     ((exp)->kind == EXP_ID || (exp)->kind == EXP_ARRAY ||                     \
      (exp)->kind == EXP_ACCESS))

#define ast_exp_add                 array_add
#define ast_exp_merge               array_merge

#ifndef _AST_EXP_T
#define _AST_EXP_T
typedef struct ast_exp_s ast_exp_t;
#endif /* ! _AST_EXP_T */

#ifndef _AST_ID_T
#define _AST_ID_T
typedef struct ast_id_s ast_id_t;
#endif /* ! _AST_ID_T */

typedef enum exp_kind_e {
    EXP_NULL        = 0,
    EXP_LIT,
    EXP_TYPE,
    EXP_ID,
    EXP_ARRAY,
    EXP_OP,
    EXP_ACCESS,
    EXP_CALL,
    EXP_SQL,
    EXP_TERNARY,
    EXP_TUPLE,
    EXP_MAX
} exp_kind_t;

typedef enum op_kind_e {
    OP_ASSIGN       = 0,
    OP_ADD,
    OP_SUB,
    OP_MUL,
    OP_DIV,
    OP_MOD,
    OP_AND,
    OP_OR,
    OP_BIT_AND,
    OP_BIT_OR,
    OP_BIT_XOR,
    OP_EQ,
    OP_NE,
    OP_LT,
    OP_GT,
    OP_LE,
    OP_GE,
    OP_RSHIFT,
    OP_LSHIFT,
    OP_INC,
    OP_DEC,
    OP_NOT,
    OP_MAX
} op_kind_t;

typedef enum sql_kind_e {
    SQL_QUERY       = 0,
    SQL_INSERT,
    SQL_UPDATE,
    SQL_DELETE,
    SQL_MAX
} sql_kind_t;

// null, true, false, 1, 1.0, 0x1, "..."
typedef struct exp_lit_s {
    ast_val_t val;
} exp_lit_t;

// primitive, struct, map
typedef struct exp_type_s {
    type_t type;
    char *name;
    ast_exp_t *k_exp;
    ast_exp_t *v_exp;
} exp_type_t;

// name
typedef struct exp_id_s {
    char *name;
} exp_id_t;

// id[idx]
typedef struct exp_array_s {
    ast_exp_t *id_exp;
    ast_exp_t *idx_exp;
} exp_array_t;

// id(param, ...)
typedef struct exp_call_s {
    ast_exp_t *id_exp;
    array_t *param_exps;
} exp_call_t;

// id.fld
typedef struct exp_access_s {
    ast_exp_t *id_exp;
    ast_exp_t *fld_exp;
} exp_access_t;

// l kind r
typedef struct exp_op_s {
    op_kind_t kind;
    ast_exp_t *l_exp;
    ast_exp_t *r_exp;
} exp_op_t;

// prefix ? infix : postfix
typedef struct exp_ternary_s {
    ast_exp_t *pre_exp;
    ast_exp_t *in_exp;
    ast_exp_t *post_exp;
} exp_ternary_t;

// dml, query
typedef struct exp_sql_s {
    sql_kind_t kind;
    char *sql;
} exp_sql_t;

// (exp, exp, exp, ...)
typedef struct exp_tuple_s {
    array_t *exps;
} exp_tuple_t;

struct ast_exp_s {
    AST_NODE_DECL;

    exp_kind_t kind;

    union {
        exp_lit_t u_lit;
        exp_type_t u_type;
        exp_id_t u_id;
        exp_array_t u_arr;
        exp_call_t u_call;
        exp_access_t u_acc;
        exp_op_t u_op;
        exp_ternary_t u_tern;
        exp_sql_t u_sql;
        exp_tuple_t u_tup;
    };

    // results of semantic checker
    // might be a part of ir_exp_t
    ast_meta_t meta;
    ast_id_t *id;
};

ast_exp_t *ast_exp_new(exp_kind_t kind, errpos_t *pos);

ast_exp_t *exp_null_new(errpos_t *pos);
ast_exp_t *exp_lit_new(errpos_t *pos);
ast_exp_t *exp_type_new(type_t type, char *name, ast_exp_t *k_exp,
                        ast_exp_t *v_exp, errpos_t *pos);
ast_exp_t *exp_id_new(char *name, errpos_t *pos);
ast_exp_t *exp_array_new(ast_exp_t *id_exp, ast_exp_t *idx_exp,
                         errpos_t *pos);
ast_exp_t *exp_call_new(ast_exp_t *id_exp, array_t *param_exps, errpos_t *pos);
ast_exp_t *exp_access_new(ast_exp_t *id_exp, ast_exp_t *fld_exp,
                          errpos_t *pos);
ast_exp_t *exp_op_new(op_kind_t kind, ast_exp_t *l_exp, ast_exp_t *r_exp,
                      errpos_t *pos);
ast_exp_t *exp_ternary_new(ast_exp_t *pre_exp, ast_exp_t *in_exp,
                           ast_exp_t *post_exp, errpos_t *pos);
ast_exp_t *exp_sql_new(sql_kind_t kind, char *sql, errpos_t *pos);
ast_exp_t *exp_tuple_new(ast_exp_t *exp, errpos_t *pos);

void ast_exp_dump(ast_exp_t *exp, int indent);

#endif /* ! _AST_EXP_H */